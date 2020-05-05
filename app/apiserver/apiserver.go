package apiserver

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	//	"github.com/pprisn/grpc_rest_mnf/app/store"
	//	"github.com/pprisn/http_rest_api/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"runtime"
	//	"github.com/go-pg/pg"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	api "github.com/pprisn/grpc_rest_mnf/api/mnf/v1"
	mnf "github.com/pprisn/grpc_rest_mnf/app/service/mnf"
	prt "github.com/pprisn/grpc_rest_mnf/app/service/part"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go/config"
	//	jaegerlog "github.com/uber/jaeger-client-go/log"
	//	"github.com/uber/jaeger-client-go/rpcmetrics"
	log "github.com/sirupsen/logrus"
	prometheus_metrics "github.com/uber/jaeger-lib/metrics/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	//store  *store.Store
}

type Service struct{}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error {

	db, err := newDB(s.config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	//store := sqlstore.New(db)

	if err := s.configureLogger(); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", s.config.BindGrpc)
	if err != nil {
		s.logger.Fatalf("Failed to listen: %v", s.config.BindGrpc)
		return err
	}
	s.logger.Info("starting BindGrpc tcp server")

	// Logrus
	logger := log.NewEntry(log.New())

	grpc_logrus.ReplaceGrpcLogger(logger)

	log.SetLevel(log.InfoLevel)

	// Prometheus monitoring
	metrics := prometheus_metrics.New()
	tracer, closer, err := config.Configuration{
		ServiceName: "mnf",
	}.NewTracer(
		config.Metrics(metrics),
	)

	if err != nil {
		s.logger.Fatalf("Cannot initialize Jaeger Tracer %s", err)
	}
	defer closer.Close()

	// Set GRPC Interceptors
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
	)

	// Register Part service, prometheus and HTTP service handler
	api.RegisterPartServiceServer(server, &prt.Service{DB: db})

	// Register Mnf service, prometheus and HTTP service handler
	api.RegisterMnfServiceServer(server, &mnf.Service{DB: db})

	//	grpc_prometheus.Register(server)

	// start Prometheus server
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		s.logger.Info("starting BindPrometheusHttp server")
		http.ListenAndServe(s.config.BindPrometheusHttp, mux)
	}()

	s.logger.Info("Starting gRPC service..")
	go server.Serve(lis) // start grpc server

	conn, err := grpc.Dial(s.config.BindGrpc, grpc.WithInsecure())
	if err != nil {
		return err
		//		panic("Couldn't contact grpc server")
	}

	mux := grpc_runtime.NewServeMux()
	err = api.RegisterPartServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return err
	}
	err = api.RegisterMnfServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return err
	}

	s.logger.Info("starting BindHttp server")
	return http.ListenAndServe(s.config.BindHttp, mux)
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	s.logger.SetFormatter(formatter)
	return nil

}

// Panic handler prints the stack trace when recovering from a panic.
var panicHandler = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	log.Errorf("panic recovered: %+v", string(buf))
	return status.Errorf(codes.Internal, "%s", p)
})
