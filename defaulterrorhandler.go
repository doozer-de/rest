package rest

import "net/http"

// DefaultErrorHandler is a default implementation of an Error Handler taken by the service framework.
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// NotFoundHandler is a default implementation of an NotFound Handler taken by the service framework.
func NotFoundHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusNotFound)
}
