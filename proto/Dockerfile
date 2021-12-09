FROM golang:1.16

RUN apt update && apt install -y protobuf-compiler


WORKDIR /tmp/go
WORKDIR /app

COPY ./proto /app

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

RUN protoc --go_out=/tmp/go --go-grpc_out=/tmp/go  -I /app *.proto

ENTRYPOINT rm -rf /api/* && cp /tmp/go/* /api && bash