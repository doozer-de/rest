# rest

[![Build Status](https://travis-ci.org/doozer-de/rest.svg?branch=master)](https://travis-ci.org/doozer-de/rest)
[![GoDoc](https://godoc.org/github.com/doozer-de/rest?status.svg)](https://godoc.org/github.com/doozer-de/rest)

Package `rest` provides functionality to:

- initiate a REST service,
- register REST routes and handler,
- handle REST requests and
- handles CORS header.

## Installation

	go get github.com/doozer-de/rest

## Usage

This library works well with [restgen](https://github.com/doozer-de/restgen).

## Limitations

Due to usage of `golang.org/x/net/context` in current Protobuf implementation
`func(ctx context.Context, w http.ResponseHandler, r *http.Request)`
signature is used for HTTP handlers instead for http.HandlerFunc.

## Credits

Parts of HTTP router: Julien Schmidt [github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
Servier based on: Christoph Seufert [github.com/0x434d53/service](https://github.com/0x434d53/service).
CORS based on: Jaana Burcu Dogan [https://github.com/martini-contrib/cors](https://github.com/martini-contrib/cors)

## License

Copyright © 2016-2017 DOOZER REAL ESTATE SYSTEMS GMBH

Licensed under the MIT license.
