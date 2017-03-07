package models

// SeedService are the methods needed to seed the app
type SeedService interface {
	SeedUsers() error
	SeedAdmins() error
}
