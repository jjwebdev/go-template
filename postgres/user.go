package postgres

import (
	"github.com/jjwebdev/go-template/models"
	"github.com/jmoiron/sqlx"
)

// UserService contains all the methods for Users
type UserService struct {
	*sqlx.DB
}

// Create will persist a user into the DB
func (us *UserService) Create(user *models.User) {
	_, err := us.DB.NamedExec(
		`
INSERT INTO users
	(first_name, last_name, email, password, session_token, data, role_id)
VALUES
	(:first_name, :last_name, :email, :password, :session_token, :data, :role_id)
`, user)
	if err != nil {
		panic(err)
	}
}

// CreateRole will persist a role into the DB
func (us *UserService) CreateRole(role *models.Role) {
	_, err := us.DB.NamedExec(
		`
INSERT INTO roles
	(name, data)
VALUES
	(:name, :data)
`, role)
	if err != nil {
		panic(err)
	}
}

// Get will directly fetch a User from the DB
func (us *UserService) Get(id int) *models.User {
	result := &models.User{}
	err := us.DB.Get(result, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	return result
}

// All will return all users from the DB
func (us *UserService) All() []*models.User {
	result := []*models.User{}
	err := us.Select(&result, "SELECT * FROM users;")
	if err != nil {
		panic(err)
	}
	return result
}
