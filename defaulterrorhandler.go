package rest

import (
	"net/http"

	"golang.org/x/net/context"
)

// DefaultErrorHandler is a default implementation of an Error Handler taken by the service framework.
func DefaultErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// NotFoundHandler is a default implementation of an NotFound Handler taken by the service framework.
func NotFoundHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusNotFound)
}
