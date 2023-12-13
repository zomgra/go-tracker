package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	bloomfilter "github.com/zomgra/bitbucket/internal/bloom"
	"github.com/zomgra/bitbucket/internal/shipment"
)

func main() {
	r := mux.NewRouter()
	shipment.AddRoutes(r)

	godotenv.Load("app.env")
	// connection, err := db.NewConnection()

	go func() {
		bloomfilter.Repository.InjectFromDB()
	}()

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8000", r))
	}()

	select {}

}
