package postgres

import (
	"sync"

	"log"

	"github.com/jjwebdev/go-template/models"
	"github.com/jmoiron/sqlx"
	"github.com/syscrusher/fake"
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

// SeedRoles will seed the roles
func (us *UserService) SeedRoles() {
	tx := us.DB.MustBegin()
	tx.MustExec("INSERT INTO roles(name) VALUES('member')")
	tx.MustExec("INSERT INTO roles(name) VALUES('admin')")
	tx.Commit()
}

// SeedUser will seed num numbers of users
func (us *UserService) SeedUser(num int, isAdmin bool) {
	log.Println("Seeding Users")
	roleID := 2
	if isAdmin {
		roleID = 1
	}
	log.Println(roleID)
	wg := &sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			u := &models.User{
				FirstName:    fake.FirstName(),
				LastName:     fake.LastName(),
				Email:        fake.EmailAddress(),
				Password:     fake.SimplePassword(),
				SessionToken: fake.FirstName(),
				RoleID:       roleID,
			}
			us.Create(u)
			wg.Done()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
