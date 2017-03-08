package http

import (
	"net/http"

	"github.com/jjwebdev/go-template/models"
	"github.com/rs/cors"
)

// Server represents an HTTP server.
type Server struct {
	Handler *Handler
	Addr    string
}

// Open opens a socket and serves the HTTP server.
func (s *Server) Open() error {
	CORSOptions := cors.Options{AllowedHeaders: []string{"X-API"}}
	CORSMux := cors.New(CORSOptions).Handler(s.Handler)
	return http.ListenAndServe(s.Addr, CORSMux)
}

// Close closes the socket.
func (s *Server) Close() error {
	panic("To be implemented")
}

// NewServer returns a new server
func NewServer(port string, userService models.UserService) *Server {
	return &Server{
		Addr: port,
		Handler: &Handler{
			UserHandler: NewUserHandler(userService),
			AppHandler:  NewAppHandler(),
			FilesMux:    NewFileServer(),
		},
	}
}
