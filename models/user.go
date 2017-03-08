package models

import "github.com/jmoiron/sqlx/types"

// User contains the columns for the users of the web application.
type User struct {
	ID           int            `json:"id" db:"id"`
	FirstName    string         `json:"firstName" db:"first_name"`
	LastName     string         `json:"lastName" db:"last_name"`
	Email        string         `json:"email" db:"email"`
	Password     string         `json:"password" db:"password"`
	SessionToken string         `json:"sessionToken" db:"session_token"`
	Data         types.JSONText `json:"data" db:"data"`
	RoleID       int            `json:"role_id" db:"role_id"`
}

// Role is a user's role
type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// UserService contains the methods that can be done with a User
type UserService interface {
	CreateRole(*Role)
	Create(*User)
	// Get(int) (*User, error)
	All() []*User
}
