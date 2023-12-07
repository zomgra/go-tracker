package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bits-and-blooms/bloom/v3"
)

func InsertShipment(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`
	_, err := Instance.db.Exec(query, barcode)
	if err != nil {
		log.Panic("Error with insert shipment", err)
	}
	return nil
}

func SelectShipment(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1` // Limit for avoid bugs in future
	row := Instance.db.QueryRow(query, barcode)

	var foundingShipmentBarcode string
	err := row.Scan(&foundingShipmentBarcode)
	if err != nil {
		return false, err
	}
	return true, nil
}

func InjectDataTo(filter *bloom.BloomFilter) {

	var cursorName string = "get_data_cursor()"
	query := fmt.Sprintf("SELECT * FROM %s AS result;", cursorName)
	rows, err := Instance.db.Query(query)
	defer rows.Close()

	if err != nil {
		log.Fatalf("Error fetching rows: %v\n", err)
	}
	for rows.Next() {
		var barcode string
		rows.Scan(&barcode)
		log.Print(barcode)
		bytes, _ := json.Marshal(barcode)
		filter.Add(bytes)
	}
}
