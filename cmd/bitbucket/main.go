package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zomgra/bitbucket/internal/routes"
	bloomfilter "github.com/zomgra/bitbucket/internal/services/bloom"
	"github.com/zomgra/bitbucket/pkg/db"
)

func main() {
	r := routes.CreateRoute()

	godotenv.Load("app.env")

	opts := db.DbOptions{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}
	db.CreateConnection(&opts)

	go func() {
		bloomfilter.Repository.InjectFromDB()
	}()

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8000", r))
	}()

	select {}

}
