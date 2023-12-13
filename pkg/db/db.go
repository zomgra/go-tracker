package db

import (
	"errors"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

type DB struct {
	db *sqlx.DB
}

func NewConnection() (*DB, error) {
	connString := os.Getenv("CONN_STRING")
	log.Print(connString)

	if connString == "" {
		return nil, errors.New("connection string is empty. Please check enviroment")
	}

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Maybe in future use for migrate
func createMigrations(connString string) {
	m, err := migrate.New("file://pkg/db/migrations", connString)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("error with migration: ", err)
	}
	m.Up()
}
