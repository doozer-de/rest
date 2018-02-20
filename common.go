package rest

import (
	"net/http"
	"sort"
)

// Middleware is an interface to concatenate functions to chains.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ErrorHandler is an interface to provide error handling.
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

// Register wraps a method, a path and a handler.
type Register struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
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
