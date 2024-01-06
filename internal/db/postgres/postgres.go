package postgres

import (
	"errors"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"github.com/zomgra/tracker/pkg/config"
)

type Client struct {
	Db *sqlx.DB
}

func NewClient(config config.PostgresConfig) (*Client, error) {
	db, err := newDB(config)
	if err != nil {
		return nil, err
	}
	client := &Client{Db: db}
	return client, nil
}

func newDB(config config.PostgresConfig) (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", config.ConnectionUrl)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	return db, nil
}

func (c *Client) Insert(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`

	_, err := c.Db.Exec(query, barcode)
	if err != nil {
		return errors.New(fmt.Sprintf("Error with insert shipment: %s", err.Error()))
	}
	return nil
}

func (c *Client) Exists(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1`

	row := c.Db.QueryRow(query, barcode)

	var foundingShipmentBarcode string
	err := row.Scan(&foundingShipmentBarcode)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) InjectDataTo(ch chan any) error {

	tx, err := c.Db.Begin()
	query := "DECLARE cursor_shipment CURSOR FOR SELECT * FROM shipments;"
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	rows, err := tx.Query("FETCH ALL FROM cursor_shipment;")
	if err != nil {
		return err
	}
	for rows.Next() {
		var barcode string

		if err := rows.Scan(&barcode); err != nil {
			c.Close()
			return err
		}
		ch <- barcode
	}

	c.Close()
	close(ch)
	return nil
}

func (c *Client) Close() error {
	return c.Db.Close()
}
