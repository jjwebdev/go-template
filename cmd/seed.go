package cmd

import (
	"github.com/jjwebdev/go-template/postgres"
)

// SeedCommand contains methods for DB migration
type SeedCommand struct {
}

// Run will begin database migration
func (mc *SeedCommand) Run(args []string) {
	db := postgres.Open("gotemplate", "develop", "develop", "localhost", "5432")

	seedUsers(db.UserService)
	defer db.Close()
}

// Usage returns the help text for the seed command
func (mc *SeedCommand) Usage() string {
	return "Seeds the database"
}

func seedRoles() {

}

func seedAdmins() {

}
func seedUsers(us *postgres.UserService) {
	us.SeedRoles()
	us.SeedUser(3, true)
	us.SeedUser(10, false)
}
