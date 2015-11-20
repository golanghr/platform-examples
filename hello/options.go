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
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/utils"
)

var (
	serviceOptions = map[string]interface{}{
		"service-name":        utils.GetFromEnvOr("HELLO_SERVICE_NAME", "Hello World"),
		"service-description": utils.GetFromEnvOr("HELLO_SERVICE_DESCRIPTION", "Basic Golang.Hr Platform hello world example service written for easier understanding of how to run and work with platform"),
		"service-version":     getFloat(utils.GetFromEnvOr("HELLO_SERVICE_VERSION", "0.1")),
		"formatter":           "text",
		"level":               logrus.DebugLevel,
		"grpc-listen-forever": getBool(utils.GetFromEnvOr("HELLO_SERVICE_GRPC_LISTEN_FOREVER", "true")),
		"grpc-addr":           utils.GetFromEnvOr("HELLO_SERVICE_GRPC_ADDR", ":4778"),
		"grpc-tls":            getBool(utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS", "true")),
		"grpc-tls-cert":       utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_CERT", "../test_data/server.crt"),
		"grpc-tls-key":        utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_KEY", "../test_data/server.key"),
	}
)

func getBool(env string) bool {
	bval, _ := strconv.ParseBool(env)
	return bval
}

func getFloat(env string) float64 {
	f, _ := strconv.ParseFloat(env, 64)
	return f
}
