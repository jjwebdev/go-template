package http

import (
	"net/http"

	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

// AppHandler represents an HTTP API handler for dials.
type AppHandler struct {
}

// NewAppHandler returns a new instance of AppHandler.
func NewAppHandler() *AppHandler {
	h := &AppHandler{}
	return h
}

// ReverseProxy
func (h *AppHandler) ReverseProxy(w http.ResponseWriter, r *http.Request) {
	r.URL = testutils.ParseURI("http://localhost:3000")
	fwd, err := forward.New()
	if err != nil {
		panic(err)
	}
	fwd.ServeHTTP(w, r)
}
