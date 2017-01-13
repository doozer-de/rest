package rest

import (
	"net/http"
	"sort"

	"golang.org/x/net/context"
)

// ContextHandler provides an interface to read and/or manipulates the context.
type ContextHandler func(context.Context, http.ResponseWriter, *http.Request)

// Middleware is an interface to concatenate functions to chains.
type Middleware func(ContextHandler) ContextHandler

// ServeWithContext handles a request withoutd a given context.
func (h ContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(nil, w, r)
}

// ServeWithContext handles a request with a given context.
func (h ContextHandler) ServeWithContext(c context.Context, w http.ResponseWriter, r *http.Request) {
	h(c, w, r)
}

// ToContextHandler converts a given handler into a ContextHandler.
func ToContextHandler(f interface{}) ContextHandler {
	var h ContextHandler

	switch f.(type) {
	case func(context.Context, http.ResponseWriter, *http.Request):
		h = ContextHandler(f.(func(context.Context, http.ResponseWriter, *http.Request)))
	case ContextHandler:
		h = f.(ContextHandler)
	case func(http.ResponseWriter, *http.Request):
		h = func(c context.Context, w http.ResponseWriter, r *http.Request) {
			f.(func(http.ResponseWriter, *http.Request))(w, r)
		}
	default:
		if h, ok := f.(http.Handler); ok {
			return ToContextHandler(h.ServeHTTP)
		}
		panic("Unsupported Handler")
	}

	return h
}

// ErrorHandler is an interface to provide error handling.
type ErrorHandler func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)

// Register wraps a method, a path and a handler.
type Register struct {
	Method  string
	Path    string
	Handler func(context.Context, http.ResponseWriter, *http.Request)
}

// HandlerRegistration provides methods neccessary to register routes and handlers.
type HandlerRegistration interface {
	GetBaseURI() string
	GetHandlersToRegister() []Register
	SetErrorHandler(h ErrorHandler) error
}

// Param wraps a key/value pair.
type Param struct {
	Key   string
	Value string
}

// Params represents a slice of Param.
type Params []Param

// Get gets the value for the given key.
func (p *Params) Get(key string) string {
	for _, pm := range *p {
		if key == pm.Key {
			return pm.Value
		}
	}

	return ""
}

// Statuser is an interface to get the status from an object.
type Statuser interface {
	Status() int
}

// SetStatus sets the HTTP status in w if v satifies the Statuser interface.
func SetStatus(w http.ResponseWriter, v interface{}) {
	if s, ok := v.(Statuser); ok {
		w.WriteHeader(s.Status())
	}
}

// Int32Slice represents a slice of int32.
type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts an int32 slice.
func (p Int32Slice) Sort() { sort.Sort(p) }
