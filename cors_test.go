package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type HttpHeaderGuardRecorder struct {
	*httptest.ResponseRecorder
	savedHeaderMap http.Header
}

func NewRecorder() *HttpHeaderGuardRecorder {
	return &HttpHeaderGuardRecorder{httptest.NewRecorder(), nil}
}

func (gr *HttpHeaderGuardRecorder) WriteHeader(code int) {
	gr.ResponseRecorder.WriteHeader(code)
	gr.savedHeaderMap = gr.ResponseRecorder.Header()
}

func (gr *HttpHeaderGuardRecorder) Header() http.Header {
	if gr.savedHeaderMap != nil {
		clone := make(http.Header)
		for k, v := range gr.savedHeaderMap {
			clone[k] = v
		}
		return clone
	} else {
		return gr.ResponseRecorder.Header()
	}
}

func AHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return
}

func TestAllowAll(t *testing.T) {
	recorder := httptest.NewRecorder()

	s := New(RestConfiguration{CORS: true}, nil)
	r, _ := http.NewRequest("GET", "path", nil)

	s.ServeHTTP(recorder, r)

	if recorder.HeaderMap.Get(headerAllowOrigin) != "*" {
		t.Errorf("Allow-Origin header should be *")
	}
}

func TestAllowRegexMatch(t *testing.T) {
	recorder := httptest.NewRecorder()
	opt := &CORSOptions{
		AllowOrigins: []string{"https://abc.org", "https://*.cs.com"},
	}

	s := New(RestConfiguration{CORS: true, CORSOptions: opt}, nil)

	origin := "https://bar.cs.com"
	r, _ := http.NewRequest("GET", "cs", nil)
	r.Header.Add("Origin", origin)
	s.ServeHTTP(recorder, r)

	headerValue := recorder.HeaderMap.Get(headerAllowOrigin)
	if headerValue != origin {
		t.Errorf("Allow-Origin header should be %v, found %v", origin, headerValue)
	}
}

func TestAllowRegexNoMatch(t *testing.T) {
	recorder := httptest.NewRecorder()

	opt := &CORSOptions{
		AllowOrigins: []string{"https://*.cs.com"},
	}

	s := New(RestConfiguration{CORS: true, CORSOptions: opt}, nil)

	origin := "https://ww.boese.com.evil.com"
	r, _ := http.NewRequest("PUT", "cs", nil)
	r.Header.Add("Origin", origin)
	s.ServeHTTP(recorder, r)

	headerValue := recorder.HeaderMap.Get(headerAllowOrigin)
	if headerValue != "" {
		t.Errorf("Allow-Origin header should not exist, found %v", headerValue)
	}
}

func TestHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()
	opt := &CORSOptions{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"PATCH", "GET"},
		AllowHeaders:     []string{"Origin", "X-whatever"},
		ExposeHeaders:    []string{"Content-Length", "Hello"},
		MaxAge:           5 * time.Minute,
	}

	s := New(RestConfiguration{CORS: true, CORSOptions: opt}, nil)

	r, _ := http.NewRequest("PUT", "foo", nil)
	s.ServeHTTP(recorder, r)

	credentialsVal := recorder.HeaderMap.Get(headerAllowCredentials)
	methodsVal := recorder.HeaderMap.Get(headerAllowMethods)
	headersVal := recorder.HeaderMap.Get(headerAllowHeaders)
	exposedHeadersVal := recorder.HeaderMap.Get(headerExposeHeaders)
	maxAgeVal := recorder.HeaderMap.Get(headerMaxAge)

	if credentialsVal != "true" {
		t.Errorf("Allow-Credentials is expected to be true, found %v", credentialsVal)
	}

	if methodsVal != "PATCH,GET" {
		t.Errorf("Allow-Methods is expected to be PATCH,GET; found %v", methodsVal)
	}

	if headersVal != "Origin,X-whatever" {
		t.Errorf("Allow-Headers is expected to be Origin,X-whatever; found %v", headersVal)
	}

	if exposedHeadersVal != "Content-Length,Hello" {
		t.Errorf("Exposef-Headers are expected to be Content-Length,Hello. Found %v", exposedHeadersVal)
	}

	if maxAgeVal != "300" {
		t.Errorf("Max-Age is expected to be 300, found %v", maxAgeVal)
	}
}

func TestDefaultAllowHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()
	opt := &CORSOptions{
		AllowAllOrigins: true,
	}

	s := New(RestConfiguration{CORS: true, CORSOptions: opt}, nil)

	r, _ := http.NewRequest("PUT", "foo", nil)
	s.ServeHTTP(recorder, r)

	headersVal := recorder.HeaderMap.Get(headerAllowHeaders)
	if headersVal != "Origin,Accept,Content-Type,Authorization" {
		t.Errorf("Allow-Headers is expected to be Origin,Accept,Content-Type,Authorization; found %v", headersVal)
	}
}

func TestPreflight(t *testing.T) {
	recorder := NewRecorder()
	opt := &CORSOptions{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH"},
		AllowHeaders:    []string{"Origin", "X-whatever", "X-CaseSensitive"},
	}

	s := New(RestConfiguration{CORS: true, CORSOptions: opt}, nil)

	s.Options("foo", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	r, _ := http.NewRequest("OPTIONS", "foo", nil)
	r.Header.Add(headerRequestMethod, "PUT")
	r.Header.Add(headerRequestHeaders, "X-whatever, x-casesensitive")
	s.ServeHTTP(recorder, r)

	headers := recorder.Header()
	methodsVal := headers.Get(headerAllowMethods)
	headersVal := headers.Get(headerAllowHeaders)
	originVal := headers.Get(headerAllowOrigin)

	if methodsVal != "PUT,PATCH" {
		t.Errorf("Allow-Methods is expected to be PUT,PATCH, found %v", methodsVal)
	}

	if !strings.Contains(headersVal, "X-whatever") {
		t.Errorf("Allow-Headers is expected to contain X-whatever, found %v", headersVal)
	}

	if !strings.Contains(headersVal, "x-casesensitive") {
		t.Errorf("Allow-Headers is expected to contain x-casesensitive, found %v", headersVal)
	}

	if originVal != "*" {
		t.Errorf("Allow-Origin is expected to be *, found %v", originVal)
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("Status code is expected to be 200, found %d", recorder.Code)
	}
}