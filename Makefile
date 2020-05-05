TARGET=grpc_mnf
PROTOC_INC_PATH=$(dir $(shell which protoc))/../include
all: clean proto build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) 	./cmd/grpc_mnf/main.go
proto:
	protoc -I $(PROTOC_INC_PATH) -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		api/mnf/v1/mnf.proto
	protoc -I $(PROTOC_INC_PATH) -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		api/mnf/v1/mnf.proto
	protoc -I $(PROTOC_INC_PATH) -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=allow_merge=true,merge_file_name=api,logtostderr=true:. \
		api/mnf/v1/mnf.proto
