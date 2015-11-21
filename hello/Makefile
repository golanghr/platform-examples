NAME = hello

GO ?= go
PROTOC ?= protoc

PROTOCFLAGS = --go_out=plugins=grpc:.
PROTOCGATEWAYFLAGS = --grpc-gateway_out=logtostderr=true:.
PROTOSPATH = protos
PLATFORMPROTOPATH = $(GOPATH)/src
GOOGLEAPIPATH = $(GOPATH)/src/github.com/gengo/grpc-gateway/third_party/googleapis
PROTOS = protos/*.proto

all: build

build:
	@echo "Building $(NAME) protocol buffers..."
	$(PROTOC) -I/usr/local/include -I. -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) -I $(GOOGLEAPIPATH) $(PROTOS) $(PROTOCFLAGS)
	$(PROTOC) -I/usr/local/include -I. -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) -I $(GOOGLEAPIPATH) $(PROTOS) $(PROTOCGATEWAYFLAGS)

test:
	$(GO) test

run:
	@echo "Building $(NAME) and running service..."
	$(GO) build && ./$(NAME)
