package postgres

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/jmoiron/sqlx"
)

type PostgresDBHandler struct {
	db *sqlx.DB
}

func NewPostgresDBHandler() (*PostgresDBHandler, error) { // Realize catching error hier
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
	return &PostgresDBHandler{db}, nil
}

func InsertShipment(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`
	connection, err := NewPostgresDBHandler()
	_, err = connection.db.Exec(query, barcode)
	if err != nil {
		log.Panic("Error with insert shipment", err)
	}
	return nil
}

func ExistShipment(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1` // Limit for avoid bugs in future
	connection, err := NewPostgresDBHandler()
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
