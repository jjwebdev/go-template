package http

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	jjwebdev "github.com/jjwebdev/go-template/models"
)

// Run begins the server
func Run(addr string, us jjwebdev.UserService) {
	routes := routes(us)

	log.Println("Your server is listening on port: 8080")

	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{"X-API"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
	}).Handler(routes)

	log.Fatalln(http.ListenAndServe(":8080", corsHandler))
}
