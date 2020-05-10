FROM golang:latest
ENV APP_NAME grpc_mnf
RUN apt-get update && apt-get install make bash
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN make build
CMD ./${APP_NAME}
EXPOSE 8080 2338 8081 5775 
