package storm

import (
	"log"

	"github.com/asdine/storm"
	jjwebdev "github.com/jjwebdev/go-template/models"
)

var db *Service

// Service is the exported DB instance
type Service struct {
	DB *storm.DB
	*UserService
}

// UserService will implement jjwebdev.UserService
type UserService struct {
	DB *storm.DB
}

// Create will create a user
func (us *UserService) Create(*jjwebdev.User) error {
	log.Println("RUNNING STORM CREATE USER")
	return nil
}

// Open will open a DB connection
func Open() *Service {
	if db == nil {
		conn, err := storm.Open("./data.db")
		if err != nil {
			log.Fatalln(err)
		}
		db = &Service{
			DB:          conn,
			UserService: &UserService{DB: conn},
		}
	}

	return db
}

// Close will close the DB connection
func Close() {
	if db != nil {
		db.DB.Close()
	}
}
