package cmd

import (
	"log"

	"github.com/jjwebdev/go-template/http"
	"github.com/jjwebdev/go-template/postgres"
)

// ServeCommand holds the server command
type ServeCommand struct {
	c *Subcommands
}

// Run will start the serve
func (sc *ServeCommand) Run() {
	db := postgres.Open("gotemplate", "develop", "develop", "localhost", "5432")

	defer db.DB.Close()

	server := http.NewServer(":8080", db.UserService)
	log.Println("Running server on 0.0.0.0:8080")
	log.Fatalln(server.Open())
}

// Usage will give the usage of the server
func (sc *ServeCommand) Usage() string {
	return "Begin the server"
}
