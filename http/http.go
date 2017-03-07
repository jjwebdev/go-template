package http

import (
	"net/http"

	"github.com/jjwebdev/go-template/models"
	"github.com/rs/cors"
)

// Server represents an HTTP server.
type Server struct {
	// Handler to serve.
	Handler *Handler
	// Bind address to open.
	Addr string
}

// Open opens a socket and serves the HTTP server.
func (s *Server) Open() error {
	options := cors.Options{AllowedHeaders: []string{"X-API"}}
	finalMuxer := cors.New(options).Handler(s.Handler)
	return http.ListenAndServe(s.Addr, finalMuxer)
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
			UserHandler: newUserHandler(userService),
			FilesMux:    newFileServer(),
		},
	}
}
