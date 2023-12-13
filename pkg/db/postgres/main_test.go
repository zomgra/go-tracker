package postgres

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	setup()

	os.Exit(m.Run())
}

func setup() {

	if err := godotenv.Load("../../../test.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println("Successfully loaded .env file")
	}
	if os.Getenv("CONN_STRING") == "" {
		os.Setenv("CONN_STRING", "postgresql://postgres:password@localhost:5432/postgresTestDB?sslmode=disable")
	}
}

// In future if have open resources
func teardown() {
	fmt.Println("Teardown: Cleaning up resources")
}
