package http

import (
	"net/http"

	"github.com/jjwebdev/go-template/models"
	"github.com/julienschmidt/httprouter"
)

// UserHandler represents an HTTP API handler for dials.
type UserHandler struct {
	*httprouter.Router
	UserService models.UserService
	Log         models.Log
}

// newUserHandler returns a new instance of UserHandler.
func newUserHandler(userService models.UserService) *UserHandler {
	h := &UserHandler{
		Router:      httprouter.New(),
		UserService: userService,
	}
	h.HandlerFunc("POST", "/users/create", insecureChain(h.userCreate))
	h.HandlerFunc("GET", "/users/all", insecureChain(h.userAll))
	h.NotFound = http.HandlerFunc(notFound)
	return h
}

// userCreate
func (h *UserHandler) userCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

// userAll
func (h *UserHandler) userAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}
