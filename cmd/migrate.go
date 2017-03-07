package cmd

import (
	"github.com/jjwebdev/go-template/postgres"
)

// MigrateCommand contains methods for DB migration
type MigrateCommand struct {
}

// Run will begin database migration
func (mc *MigrateCommand) Run(args []string) {
	db := postgres.Open("gotemplate", "develop", "develop", "localhost", "5432")
	db.Migrate()
	defer db.Close()
}

func (mc *MigrateCommand) Usage() string {
	return "Migrate the database"
}
