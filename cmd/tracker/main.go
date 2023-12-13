package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	bloomfilter "github.com/zomgra/tracker/internal/bloom"
	"github.com/zomgra/tracker/internal/shipment"
)

func main() {
	r := mux.NewRouter()
	shipmentRepository := shipment.NewRepository()
	bloomRepository := bloomfilter.NewRepository()
	handler := shipment.NewHandler(bloomRepository, shipmentRepository)

	shipment.AddRoutes(r, handler)

	godotenv.Load("app.env")
	// connection, err := db.NewConnection()

	go func() {
		bloomRepository.InjectFromDB()
	}()

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8000", r))
	}()

	select {}

}
