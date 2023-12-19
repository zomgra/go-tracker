package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(config Config) (*Client, error) {
	db, err := newDB(config)
	if err != nil {
		return nil, err
	}
	client := &Client{db: db}
	return client, nil
}

func newDB(config Config) (*sqlx.DB, error) { // Realize catching errors hier

	db, err := sqlx.Connect("postgres", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return db, nil
}

func (c *Client) Insert(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`

	_, err := c.db.Exec(query, barcode)
	if err != nil {
		log.Panic("Error with insert shipment", err)
	}
	return nil
}

func (c *Client) Exists(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1` // Limit for avoid bugs in future

	row := c.db.QueryRow(query, barcode)

	var foundingShipmentBarcode string
	err := row.Scan(&foundingShipmentBarcode)
	if err != nil {
		return false, err
	}

	return true, nil
}

// add returning to channel
func (c *Client) InjectDataTo(ch chan any) error {

	tx, err := c.db.Begin()
	query := "DECLARE cursor_shipment CURSOR FOR SELECT * FROM shipments;"
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}
	//error
	rows, err := tx.Query("FETCH ALL FROM cursor_shipment;")
	if err != nil {
		return err
	}
	for rows.Next() {
		var barcode string

		if err := rows.Scan(&barcode); err != nil {
			//return err
			//catch error
			c.Close()
		}
		ch <- barcode
	}

	c.Close()
	return nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
