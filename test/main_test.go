package test

import (
	"log"
	"os"
	"testing"

	"github.com/zomgra/tracker/internal/db/postgres"
	"github.com/zomgra/tracker/pkg/config"
	"github.com/zomgra/tracker/pkg/db"
)

func TestMain(m *testing.M) {
	setup()

	os.Exit(m.Run())
}

var client db.Client

func createPostgresClient() db.Client {
	dbConfig := config.PostgresConfig{}
	config.SetDbConfig(&dbConfig)
	dbClient, err := postgres.NewClient(dbConfig)

	if err != nil {
		log.Fatalf("Error creating db client: %v", err)
	}

	return dbClient
}

func setup() {

	path, err := config.FetchConfigPath()
	log.Print(path)
	if err != nil {
		log.Fatalf("Error patching config path: %v", err)
	}
	if err = config.IncludeEnv(path); err != nil {
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
