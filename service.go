// Copyright (c) 2016 Christoph Seufert, Inc.

package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

// PARAMS it the key under which the Params of the Router will be written to the context
const PARAMS = "PARAMS"

// GRPCRESTService implements a HTTP Router + Some Helpers for chaining and error handling. It is used for the GRPC-REST Gateway
type GRPCRESTService struct {
	routes          map[string]*node
	preChain        []ContextHandler
	chain           []Middleware
	notFoundHandler ErrorHandler // The Route could not be found
	errorHandler    ErrorHandler // Handle errors in general. We expected the DError and the data in the context
	trimSlash       bool
	baseURI         string
	corsEnabled     bool
	corsOptions     *CORSOptions
}

// RestConfiguration container the configuration Parameter needed to initialize the GRPCRESTService
type RestConfiguration struct {
	// BaseURI can prefix the path for this handler. A common example for REST would be the version like "/v1"
	BaseURI string
	//ErrorHandler will be called in the middlewares and the framework if an error occurs. If not present the DefaultErrorHandler will be used
	ErrorHandler ErrorHandler
	// Chain allow to execute some middlewares around the actual handler. In the slice the the most left (index 0) middleware will be executed first (from the view of an incoming request)
	Chain []Middleware // Middlewares that wrap the Route Handlers. The Route Selection happens before
	// CORS set if CORS Headers should be provided,
	CORS bool
	// CORSOptions allows to give the CORS Options to the service. If CORS is set to true and no CORSOptions are given all origins will be allowed
	CORSOptions *CORSOptions
}

// New created a new GRPCRESTServices and applies the configuration and register the handlers given by the registrators
func New(cfg RestConfiguration, registrators []HandlerRegistration) *GRPCRESTService {
	s := &GRPCRESTService{
		baseURI:         path.Join("/", cfg.BaseURI, "/"),
		routes:          map[string]*node{},
		trimSlash:       true,
		chain:           cfg.Chain,
		errorHandler:    cfg.ErrorHandler,
		notFoundHandler: cfg.ErrorHandler,
		corsEnabled:     cfg.CORS,
		corsOptions:     cfg.CORSOptions,
	}
	if s.errorHandler == nil {
		s.errorHandler = DefaultErrorHandler
		s.notFoundHandler = DefaultErrorHandler
	}

	for _, reg := range registrators {
		err := s.Register(reg.GetHandlersToRegister(), reg.GetBaseURI())
		if err != nil {
			log.Fatalf("Error registering handlers: %s", err)
		}
		err = reg.SetErrorHandler(s.errorHandler)

		if err != nil {
			log.Fatalf("Error registering handlers %s", err)
		}
	}

	if s.corsEnabled {
		o := s.corsOptions

		if o == nil {
			o = DefaultCORSOptions()
		}
		s.preChain = append(s.preChain, NewCORS(o))
	}

	return s
}

// Register registers a list of handers/paths/methods wrapping them in the middleware chain
func (s *GRPCRESTService) Register(r []Register, baseURI string) error {
	for _, r := range r {
		h := r.Handler

		for i := len(s.chain) - 1; i >= 0; i-- {
			h = s.chain[i](h)
		}

		route := path.Join(baseURI, r.Path)
		err := s.Route(r.Method, route, h)
		if err != nil {
			return err
		}
	}

	return nil
}

// ServeHTTP is there to fulfill http.Handler interface, other then that the framework uses the libs.ContextHandler as the signature
func (s *GRPCRESTService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeWithContext(context.Background(), w, r)
}

// ServeWithContext is the Entrypoint for an request.
func (s *GRPCRESTService) ServeWithContext(c context.Context, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Path != "/" && s.trimSlash {
		r.URL.Path = strings.TrimRight(r.URL.Path, "/")
	}

	for _, p := range s.preChain {
		p(c, w, r)
	}

	var h ContextHandler
	var ps Params

	n, ok := s.routes[r.Method]

	if ok {
		h, ps, _ = n.getValue(r.URL.Path)

		c = context.WithValue(c, "params", ps)
	}

	if h == nil {
		if s.notFoundHandler != nil {
			s.notFoundHandler(c, w, r, fmt.Errorf("contextHandler not found"))
		} else {
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	} else {
		c = context.WithValue(c, PARAMS, ps)

		h(c, w, r)
	}
}

// Route registers a handler for certain http method/route
func (s *GRPCRESTService) Route(method, uri string, handler ContextHandler) error {
	h := ToContextHandler(handler)

	if n := s.routes[method]; n == nil {
		s.routes[method] = &node{}
	}

	s.routes[method].addRoute(path.Join(s.baseURI, strings.TrimRight(uri, "/")), h)

	return nil
}

// Get registers a handler for GET and the given uri
func (s *GRPCRESTService) Get(uri string, handler ContextHandler) {
	s.Route("GET", uri, handler)
}

// Post registers a handler for POST and the given uri
func (s *GRPCRESTService) Post(uri string, handler ContextHandler) {
	s.Route("POST", uri, handler)
}

// Put registers a handler for PUT and the given uri
func (s *GRPCRESTService) Put(uri string, handler ContextHandler) {
	s.Route("PUT", uri, handler)
}

// Delete registers a handler for DELETE and the given uri
func (s *GRPCRESTService) Delete(uri string, handler ContextHandler) {
	s.Route("DELETE", uri, handler)
}

// Patch registers a handler for PATCH and the given uri
func (s *GRPCRESTService) Patch(uri string, handler ContextHandler) {
	s.Route("PATCH", uri, handler)
}

// Head registers a handler for HEAD and the given uri
func (s *GRPCRESTService) Head(uri string, handler ContextHandler) {
	s.Route("HEAD", uri, handler)
}

func (s *GRPCRESTService) Options(uri string, handler ContextHandler) {
	s.Route("OPTIONS", uri, handler)
}