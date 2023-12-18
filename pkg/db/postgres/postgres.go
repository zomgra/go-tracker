package postgres

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/zomgra/tracker/configs"
	"github.com/zomgra/tracker/pkg/bloomfilter"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

type PostgresDBHandler struct {
	db *sqlx.DB
}

type PostgresDBClient struct {
	handler PostgresDBHandler
}

func NewDBClient(config configs.DBConfig) (*PostgresDBClient, error) {
	handler, err := newPostgresDBHandler(config)
	if err != nil {
		return nil, err
	}
	client := &PostgresDBClient{handler: *handler}
	return client, nil
}

func newPostgresDBHandler(config configs.DBConfig) (*PostgresDBHandler, error) { // Realize catching errors hier

	db, err := sqlx.Connect("postgres", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return &PostgresDBHandler{db}, nil
}

func (c *PostgresDBClient) Insert(barcode string) error {
	query := `INSERT INTO shipments (barcode) VALUES ($1)`

	_, err := c.handler.db.Exec(query, barcode)
	if err != nil {
		log.Panic("Error with insert shipment", err)
	}
	return nil
}

func (c *PostgresDBClient) Exist(barcode string) (bool, error) {
	query := `SELECT * FROM shipments WHERE barcode = $1 LIMIT 1` // Limit for avoid bugs in future

	row := c.handler.db.QueryRow(query, barcode)

	var foundingShipmentBarcode string
	err := row.Scan(&foundingShipmentBarcode)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *PostgresDBClient) InjectDataTo(filter *bloomfilter.BloomFilterHelper) error {

	tx, err := c.handler.db.Begin()
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
		rows.Scan(&barcode)
		bytes, _ := json.Marshal(barcode)
		filter.Add(bytes)
	}
	time.Sleep(5 * time.Second)
	return nil
}

func (c *PostgresDBClient) Close() error {
	return c.handler.db.Close()
}
