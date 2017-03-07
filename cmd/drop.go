package cmd

import (
	"github.com/jjwebdev/go-template/postgres"
)

// DropCommand contains methods for DB migration
type DropCommand struct {
}

// Run will begin database migration
func (mc *DropCommand) Run(args []string) {
	db := postgres.Open("gotemplate", "develop", "develop", "localhost", "5432")
	db.Drop()
	defer db.Close()
}

func (mc *DropCommand) Usage() string {
	return "Drop the database"
}
