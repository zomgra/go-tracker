package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

var Instance *DB

type DbOptions struct {
	User, Password, Addr, Database string
}

func CreateConnection(o *DbOptions) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", o.User, o.Password, o.Addr, o.Database)
	log.Print(connString)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	createMigrations(connString)
	Instance = &DB{db}
}

func createMigrations(connString string) {
	m, err := migrate.New("file://pkg/db/migrations", connString)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("error with migration: ", err)
	}
	m.Up()
}
