package cmd

import (
	"github.com/jjwebdev/go-template/postgres"
	"github.com/jjwebdev/go-template/seed"
)

// SeedCommand contains methods for DB migration
type SeedCommand struct {
}

// Run will begin database migration
func (mc *SeedCommand) Run() {
	db := postgres.Open("gotemplate", "develop", "develop", "localhost", "5432")
	ss := seed.NewSeedService(db.UserService)
	ss.SeedRoles()
	ss.SeedUsers(2, true)
	ss.SeedUsers(10, false)
	defer db.Close()
}

// Usage returns the help text for the seed command
func (mc *SeedCommand) Usage() string {
	return "Seeds the database"
}
