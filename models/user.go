package models

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

// UserService contains the methods that can be done with a User
type UserService interface {
	Create(*User) error
	// Get(int) (*User, error)
	// All() ([]*User, error)
}
