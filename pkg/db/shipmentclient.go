package db

import (
	"encoding/json"
	"log"
	"time"

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

	tx, err := Instance.db.Begin()
	query := "DECLARE cursor_shipment CURSOR FOR SELECT * FROM shipments;"
	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("error creating cursor: ", err)
	}

	// Выполняем FETCH для извлечения данных
	rows, err := tx.Query("FETCH ALL FROM cursor_shipment;")
	if err != nil {
		log.Fatal("error fetching using cursor: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var barcode string
		rows.Scan(&barcode)
		bytes, _ := json.Marshal(barcode)
		filter.Add(bytes)
	}
	time.Sleep(5 * time.Second)
}
