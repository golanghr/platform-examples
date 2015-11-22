/*
Copyright (c) 2015 Golang Croatia
All rights reserved.

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package main ...
package main

import (
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/utils"
	. "github.com/smartystreets/goconvey/convey"

	pb "github.com/golanghr/platform-examples/hello/protos"
)

func getHelloOptions() (options.Options, error) {
	return options.New("memo", map[string]interface{}{
		"service-name":        utils.GetFromEnvOr("HELLO_SERVICE_NAME", "Hello World Test"),
		"service-description": utils.GetFromEnvOr("HELLO_SERVICE_DESCRIPTION", "Basic Golang.Hr Platform hello world example service written for easier understanding of how to run and work with platform"),
		"service-version":     getFloat(utils.GetFromEnvOr("HELLO_SERVICE_VERSION", "0.1")),
		"formatter":           "text",
		"level":               logrus.DebugLevel,
		"manager-interrupt-wait-timeout": getInt(utils.GetFromEnvOr("HELLO_SERVICE_MANAGER_INTERRUPT_TIMEOUT", "1")),
		"grpc-listen-forever":            getBool(utils.GetFromEnvOr("HELLO_SERVICE_GRPC_LISTEN_FOREVER", "true")),
		"grpc-addr":                      utils.GetFromEnvOr("HELLO_SERVICE_GRPC_ADDR", ":4771"),
		"grpc-tls":                       getBool(utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS", "true")),
		"grpc-tls-domain":                utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_DOMAIN", "golang.hr"),
		"grpc-tls-cert":                  utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_CERT", "test_data/server.crt"),
		"grpc-tls-key":                   utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_KEY", "test_data/server.key"),
		"http-addr":                      utils.GetFromEnvOr("HELLO_SERVICE_HTTP_ADDR", ":8072"),
		"http-listen-forever":            getBool(utils.GetFromEnvOr("HELLO_SERVICE_HTTP_LISTEN_FOREVER", "true")),
	})
}

func getHelloService(opts options.Options) (*Service, error) {
	return NewService(opts, logger)
}

func TestGRPCServerExample(t *testing.T) {
	opts, _ := getHelloOptions()
	service, err := getHelloService(opts)

	Convey("Should be service without any errors", t, func() {
		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)
	})

	Convey("Should receive HelloWorld platform.Response provided by HelloWorld handler", t, func() {
		go func() { service.Start() }()
		defer service.Terminate()

		address, ok := service.Options.Get("grpc-addr")
		So(ok, ShouldBeTrue)

		var gopts []grpc.DialOption

		var creds credentials.TransportAuthenticator
		domain, _ := opts.Get("grpc-tls-domain")
		creds, err = credentials.NewClientTLSFromFile("test_data/server.crt", domain.String())
		So(err, ShouldBeNil)

		gopts = append(gopts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(address.String(), gopts...)
		So(err, ShouldBeNil)
		defer conn.Close()

		client := pb.NewHelloClient(conn)

		helloResponse, err := client.HelloWorld(context.Background(), &pb.HelloRequest{})
		So(err, ShouldBeNil)
		So(helloResponse.Message, ShouldContainSubstring, "Hello From Golang.HR Micro Platform!")
	})
}

func TestHTTPServerExample(t *testing.T) {
	opts, _ := getHelloOptions()

	// We gotta update these as there are multiple tests spawning out multiple listeners
	opts.Set("http-addr", ":8071")
	opts.Set("grpc-addr", ":4772")

	service, err := getHelloService(opts)

	Convey("Should be yet another service without any errors", t, func() {
		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)
	})

	Convey("Should receive JSON HelloWorld platform.Response provided by HelloWorld handler", t, func() {
		go func() { service.Start() }()
		defer service.Terminate()

		time.Sleep(1 * time.Second)
	})
}
