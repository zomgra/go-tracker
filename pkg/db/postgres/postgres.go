package postgres

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/bits-and-blooms/bloom"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

type PostgresDBHandler struct {
	db *sqlx.DB
}

func NewPostgresDBHandler() (*PostgresDBHandler, error) { // Realize catching errors hier
	connString := os.Getenv("CONN_STRING")
	log.Print(connString)

	if connString == "" {
		log.Println("CONN STRING EMPTY")
		return nil, errors.New("connection string is empty. Please check enviroment")
	}

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return &PostgresDBHandler{db}, nil
}

func InsertShipment(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`
	connection, err := NewPostgresDBHandler()

	defer connection.db.Close()

	_, err = connection.db.Exec(query, barcode)
	if err != nil {
		log.Panic("Error with insert shipment", err)
	}
	return nil
}

func ExistShipment(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1` // Limit for avoid bugs in future
	connection, err := NewPostgresDBHandler()
	defer connection.db.Close()
	row := connection.db.QueryRow(query, barcode)

	var foundingShipmentBarcode string
	err = row.Scan(&foundingShipmentBarcode)
	if err != nil {
		return false, err
	}

	return true, nil
}

func InjectDataTo(filter *bloom.BloomFilter) error {
	connection, err := NewPostgresDBHandler()
	if err != nil {
		return err
	}
	tx, err := connection.db.Begin()
	query := "DECLARE cursor_shipment CURSOR FOR SELECT * FROM shipments;"
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	rows, err := tx.Query("FETCH ALL FROM cursor_shipment;")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var barcode string
		rows.Scan(&barcode)
		bytes, _ := json.Marshal(barcode)
		filter.Add(bytes)
	}
	time.Sleep(5 * time.Second)
	return nil
}
