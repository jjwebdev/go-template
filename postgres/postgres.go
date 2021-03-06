package postgres

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"

	// Postgres driver
	_ "github.com/lib/pq"
)

var instance *DB

// DB contains all the methods the DB can do
type DB struct {
	*UserService
}

// Open will open the postgres connection
func Open(dbname, username, password, host, port string) *DB {
	var err error
	db, err := sqlx.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbname, username, password, host, port))

	if err != nil {
		panic(err)
	}

	userService := &UserService{
		db,
	}

	instance = &DB{
		userService,
	}

	return instance
}

// Close will close the postgres connection
func Close() error {
	return instance.Close()
}

// Client will return an instance of the postgres connection
func Client() *DB {
	if instance == nil {
		panic("DB instance not initialised")
	}
	return instance
}

// Migrate begins the migration for postgres
func (db *DB) Migrate() {
	schema, err := ioutil.ReadFile("./postgres/migrate.sql")
	if err != nil {
		panic(err)
	}
	db.DB.MustExec(string(schema))
}

// Drop will drop the entire schema
func (db *DB) Drop() {
	_, err := instance.DB.Exec("DROP SCHEMA public CASCADE;")
	if err != nil {
		log.Println(err)
	}
}
