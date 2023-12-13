package db

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
)

func InsertShipment(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`
	connection, err := NewPool()
	_, err = connection.db.Exec(query, barcode)
	if err != nil {
		log.Panic("Error with insert shipment", err)
	}
	return nil
}

func ExistShipment(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1` // Limit for avoid bugs in future
	connection, err := NewPool()
	row := connection.db.QueryRow(query, barcode)

	var foundingShipmentBarcode string
	err = row.Scan(&foundingShipmentBarcode)
	if err != nil {
		return false, err
	}
	return true, nil
}

func InjectDataTo(filter *bloom.BloomFilter) error {
	connection, err := NewPool()
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
