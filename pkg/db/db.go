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

type DBHandler struct {
	db *sqlx.DB
}

func NewPool() (*DBHandler, error) { // Realize catching error hier
	connString := os.Getenv("CONN_STRING")
	log.Print(connString)

	if connString == "" {
		return nil, errors.New("connection string is empty. Please check enviroment")
	}

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return &DBHandler{db}, nil
}

// Maybe in future use for stationary migrate
func createMigrations(connString string) {
	m, err := migrate.New("file://pkg/db/migrations", connString)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("error with migration: ", err)
	}
	m.Up()
}
