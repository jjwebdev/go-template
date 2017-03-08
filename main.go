package main

import (
	"fmt"
	"os"

	"github.com/jjwebdev/go-template/cmd"
)

var helptext = `
Welcome to go-template.

Please use one of the following commands:

	serve
	migrate
	seed

Have a pleasant day!
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(helptext)
		os.Exit(1)
	}

	subCommands := cmd.New("Main", "Command")
	server := &cmd.ServeCommand{}
	drop := &cmd.DropCommand{}
	migrate := &cmd.MigrateCommand{}
	seed := &cmd.SeedCommand{}

	subCommands.Register("serve", "Begin the server", server)
	subCommands.Register("drop", "Drop the DB", drop)
	subCommands.Register("migrate", "Migrate the DB schema", migrate)
	subCommands.Register("seed", "seed the DB schema", seed)

	for _, cmd := range args {
		subCommands.Run(cmd)
	}
}
