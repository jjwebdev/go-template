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
}

// NewUserHandler returns a new instance of UserHandler.
func NewUserHandler(userService models.UserService) *UserHandler {
	h := &UserHandler{
		Router:      httprouter.New(),
		UserService: userService,
	}
	h.HandlerFunc("POST", "/users/roles/create", insecureChain(h.RoleCreate))
	h.HandlerFunc("POST", "/users/create", insecureChain(h.UserCreate))
	h.HandlerFunc("GET", "/users/all", insecureChain(h.UserAll))
	h.HandlerFunc("GET", "/users/show", insecureChain(h.UserGet))
	h.NotFound = http.HandlerFunc(notFound)
	return h
}

// UserCreate creates a User
func (h *UserHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	mustDecodeJSON(r, newUser)
	h.UserService.Create(newUser)
}

// RoleCreate creates a user role
func (h *UserHandler) RoleCreate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "RoleCreate",
		Description: "Creates a user role",
	}
	newRole := &models.Role{}
	mustDecodeJSON(r, newRole)
	h.UserService.CreateRole(newRole)
	marshalAndRespond(details, w, newRole)
}

// UserAll returns all users
func (h *UserHandler) UserAll(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "UserAll",
		Description: "Returns all users",
	}

	marshalAndRespond(details, w, h.UserService.All())
}

// UserGet returns a single user
func (h *UserHandler) UserGet(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "UserGet",
		Description: "Returns a single user",
	}

	id := mustParseInt(r, "id")
	marshalAndRespond(details, w, h.UserService.Get(id))
}
