package http

import (
	"net/http"
	"strings"
)

// Handler is a collection of all the service handlers.
type Handler struct {
	UserHandler *UserHandler
	FilesMux    *http.ServeMux
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API") != "" {
		if strings.HasPrefix(r.URL.Path, "/users") {
			h.UserHandler.ServeHTTP(w, r)
		} else {
			notFound(w, r)
		}
	} else {
		h.FilesMux.ServeHTTP(w, r)
	}
}
