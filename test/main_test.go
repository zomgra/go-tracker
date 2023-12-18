package postgres

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/zomgra/tracker/configs"
	"github.com/zomgra/tracker/pkg/db"
	"github.com/zomgra/tracker/pkg/db/postgres"
)

func TestMain(m *testing.M) {
	setup()

	os.Exit(m.Run())
}

var client db.Client

func createPostgresClient() db.Client {
	dbConfig := configs.DBConfig{ConnectionString: os.Getenv("CONNECTION_STRING")}
	log.Println(dbConfig)
	client, _ := postgres.NewDBClient(dbConfig)

	return client
}

func setup() {

	if err := godotenv.Load("../configs/test.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println("Successfully loaded .env file")
	}
	if os.Getenv("CONN_STRING") == "" {
		os.Setenv("CONN_STRING", "postgresql://postgres:password@localhost:5432/postgresTestDB?sslmode=disable")
	}
	log.Println("Creating client")
	client = createPostgresClient()
}

// In future if have open resources
func teardown() {
	fmt.Println("Teardown: Cleaning up resources")
}
