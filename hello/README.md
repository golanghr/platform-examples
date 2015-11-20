[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/golanghr/platform-examples/tree/master/LICENSE.md)
[![Build Status](https://travis-ci.org/golanghr/platform.svg)](https://travis-ci.org/golanghr/platform-examples/hello)
[![Go 1.4 Ready](https://img.shields.io/badge/Go%201.4-Ready-green.svg?style=flat)]()
[![Go 1.5 Ready](https://img.shields.io/badge/Go%201.5-Ready-green.svg?style=flat)]()

# Platform Examples - Hello World
Basic [Golang.hr Platform] example designed to show how to successfully establish [gRPC] and HTTP servers.

- [Protocol Buffers] that are required by the hello service
- [gRPC] server
- HTTP server
- HelloWorld handler accessible by WWW

### Important conceptual notes

- Each handler is designed to be [Protocol Buffers] service method.  
- By defining single handler, you are defining ALL handlers. Same means that if you define
  [gRPC] adapter (which is a must), by attaching HTTP server you immediately have access to it over HTTP, TLS, etc...
- We are and will always be compatible with [gRPC] as we want to utilize their clients so we don't have to re-invent the wheels.

**TBD - This is still in drawing mode so bare with us...**

[gRPC]: <http://www.grpc.io/>
[Protocol Buffers]: <https://developers.google.com/protocol-buffers/>

[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
[Platform Users]: <https://github.com/golanghr/platform-users>
[filing an issue]: <https://github.com/golanghr/platform/issues/new>

[Golang.hr Slack]: <http://slack.golang.hr>
[Golang.hr Facebook]: <https://www.facebook.com/groups/golanghr/>
