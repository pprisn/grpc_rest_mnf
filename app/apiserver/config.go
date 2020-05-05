package apiserver

//import "github.com/pprisn/grpc_rest_mnf/app/store"

// Config ...
type Config struct {
	BindHttp           string  `toml:"bind_http"`
	BindGrpc           string  `toml:"bind_grpc"`
	BindPrometheusHttp string  `toml:"bind_prometheus_http"`
	JaegerSampler      float64 `toml:"jaeger_sampler"`
	JaegerHost         string  `toml:"jaeger_host"`
	JaegerPort         string  `toml:"jaeger_port"`
	LogLevel           string  `toml:"log_level"`
	//	Store              *store.Config
	DatabaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindHttp:           ":8080",
		BindGrpc:           ":2338",
		BindPrometheusHttp: ":8081",
		JaegerSampler:      0.05,
		JaegerHost:         "127.0.0.1",
		JaegerPort:         ":5775",
		LogLevel:           "debug",
		//		Store:              store.NewConfig(),
	}
}
