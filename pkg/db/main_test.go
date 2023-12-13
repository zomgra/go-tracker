package db

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	if err := godotenv.Load("../../test.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println("Successfully loaded .env file")
	}

	result := m.Run()
	if result != 0 {
		log.Printf("Tests failed with code %d", result)
	}
	os.Exit(result)
}
