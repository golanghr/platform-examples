FROM debian:wheezy

ENV DEBIAN_FRONTEND noninteractive
ENV TERM dumb

RUN echo "deb http://ftp.us.debian.org/debian/ wheezy main contrib non-free" >> /etc/apt/sources.list
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

ENV HELLO_SERVICE_NAME "Hello World"
ENV HELLO_SERVICE_DESCRIPTION "Basic Golang.Hr Platform hello world example service written for easier understanding of how to run and work with platform"
ENV HELLO_SERVICE_VERSION "0.1"
ENV HELLO_SERVICE_MANAGER_INTERRUPT_TIMEOUT "5"
ENV HELLO_SERVICE_GRPC_LISTEN_FOREVER "true"
ENV HELLO_SERVICE_GRPC_ADDR ":4780"
ENV HELLO_SERVICE_GRPC_TLS "true"
ENV HELLO_SERVICE_GRPC_TLS_DOMAIN "golang.hr"

ADD test_data/server.crt /certs/server.crt
ADD test_data/server.key /certs/server.key
ENV HELLO_SERVICE_GRPC_TLS_CERT "/certs/server.crt"
ENV HELLO_SERVICE_GRPC_TLS_KEY "/certs/server.key"

ENV HELLO_SERVICE_HTTP_ADDR ":8010"
ENV HELLO_SERVICE_HTTP_LISTEN_FOREVER "true"

EXPOSE 4780
EXPOSE 8010

ADD build/platform-examples-hello /usr/bin/platform-examples-hello
ENTRYPOINT ["platform-examples-hello"]
