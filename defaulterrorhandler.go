package rest

import (
	"net/http"

	"golang.org/x/net/context"
)

// DefaultErrorHandler is a default implementation of an Error Handler taken by the service framework
func DefaultErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func NotFoundHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}
