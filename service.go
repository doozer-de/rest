// Copyright (c) 2016 Christoph Seufert

package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

// paramsKey defines context paramteters key
type paramsKey struct{}

// Service implements a HTTP Router + Some Helpers for chaining and error handling. It is used for the GRPC-REST Gateway
type Service struct {
	routes          map[string]*node
	optionsChain    []http.Handler
	chain           []Middleware
	notFoundHandler ErrorHandler // The Route could not be found
	errorHandler    ErrorHandler // Handle errors in general. We expected the DError and the data in the context
	trimSlash       bool
	baseURI         string
	corsEnabled     bool
	corsOptions     *CORSOptions
}

// Configuration container the configuration Parameter needed to initialize the GRPCRESTService
type Configuration struct {
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
func New(cfg Configuration, registrators []HandlerRegistration) *Service {
	s := &Service{
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
		s.optionsChain = append(s.optionsChain, NewCORS(o))
	}

	return s
}

// Register registers a list of handers/paths/methods wrapping them in the middleware chain
func (s *Service) Register(r []Register, baseURI string) error {
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

// ServeHTTP is the Entrypoint for an request.
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()

	if r.URL.Path != "/" && s.trimSlash {
		r.URL.Path = strings.TrimRight(r.URL.Path, "/")
	}

	for _, p := range s.optionsChain {
		p.ServeHTTP(w, r)
	}
	if r.Method == http.MethodOptions {
		return
	}

	var h http.Handler
	var ps Params

	n, ok := s.routes[r.Method]

	if ok {
		h, ps, _ = n.getValue(r.URL.Path)
		ctx = context.WithValue(ctx, paramsKey{}, ps)
	}

	if h == nil {
		if s.notFoundHandler != nil {
			s.notFoundHandler(w, r, fmt.Errorf("contextHandler not found"))
		} else {
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	} else {
		ctx = context.WithValue(ctx, paramsKey{}, ps)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

func GetParams(ctx context.Context) Params {
	return ctx.Value(paramsKey{}).(Params)
}

// Route registers a handler for certain http method/route
func (s *Service) Route(method, uri string, handler http.Handler) error {
	if n := s.routes[method]; n == nil {
		s.routes[method] = &node{}
	}

	s.routes[method].addRoute(path.Join(s.baseURI, strings.TrimRight(uri, "/")), handler)

	return nil
}

// Get registers a handler for GET and the given uri
func (s *Service) Get(uri string, handler http.Handler) {
	s.Route(http.MethodGet, uri, handler)
}

// Post registers a handler for POST and the given uri
func (s *Service) Post(uri string, handler http.Handler) {
	s.Route(http.MethodPost, uri, handler)
}

// Put registers a handler for PUT and the given uri
func (s *Service) Put(uri string, handler http.Handler) {
	s.Route(http.MethodPut, uri, handler)
}

// Delete registers a handler for DELETE and the given uri
func (s *Service) Delete(uri string, handler http.Handler) {
	s.Route(http.MethodDelete, uri, handler)
}

// Patch registers a handler for PATCH and the given uri
func (s *Service) Patch(uri string, handler http.Handler) {
	s.Route(http.MethodPatch, uri, handler)
}

// Head registers a handler for HEAD and the given uri
func (s *Service) Head(uri string, handler http.Handler) {
	s.Route(http.MethodHead, uri, handler)
}

// Options registers a handler for OPTIONS and the given uri
func (s *Service) Options(uri string, handler http.Handler) {
	s.Route(http.MethodOptions, uri, handler)
}
