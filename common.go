package rest

import (
	"net/http"
	"sort"

	"golang.org/x/net/context"
)

type ContextHandler func(context.Context, http.ResponseWriter, *http.Request)

type Middleware func(ContextHandler) ContextHandler

func (h ContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(nil, w, r)
}

func (h ContextHandler) ServeWithContext(c context.Context, w http.ResponseWriter, r *http.Request) {
	h(c, w, r)
}

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

type ErrorHandler func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)

type Register struct {
	Method  string
	Path    string
	Handler func(context.Context, http.ResponseWriter, *http.Request)
}

type HandlerRegistration interface {
	GetBaseURI() string
	GetHandlersToRegister() []Register
	SetErrorHandler(h ErrorHandler) error
}

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (p *Params) Get(key string) string {
	for _, pm := range *p {
		if key == pm.Key {
			return pm.Value
		}
	}

	return ""
}

type Statuser interface {
	Status() int
}

func SetStatus(w http.ResponseWriter, v interface{}) {
	if s, ok := v.(Statuser); ok {
		w.WriteHeader(s.Status())
	}
}

type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int32Slice) Sort()              { sort.Sort(p) }
