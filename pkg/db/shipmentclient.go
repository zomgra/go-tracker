package db

import "log"

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
	var foundingShipment string
	err := row.Scan(&foundingShipment)
	if err != nil {
		return false, err
	}
	return true, nil
}
