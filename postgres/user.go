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

func (us *UserService) CreateRole(role *models.Role) {
	_, err := us.DB.NamedExec(
		`
INSERT INTO roles
	(name)
VALUES
	(:name)
`, role)
	if err != nil {
		panic(err)
	}
}
