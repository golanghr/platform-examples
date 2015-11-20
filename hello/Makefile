NAME = hello

GO ?= go
PROTOC ?= protoc

PROTOCFLAGS = --go_out=plugins=grpc:.
PROTOSPATH = protos
PLATFORMPROTOPATH = $(GOPATH)/src
PROTOS = protos/*.proto

all: build

build:
	@echo "Building $(NAME) protocol buffers..."
	$(PROTOC) -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) $(PROTOS) $(PROTOCFLAGS)

test:
	$(GO) test

run:
	@echo "Building $(NAME) and running service..."
	$(GO) build && ./$(NAME)
