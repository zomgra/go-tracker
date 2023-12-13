package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/zomgra/tracker/internal/app"
	"github.com/zomgra/tracker/internal/shipment"
)

func main() {
	r := mux.NewRouter()
	shipmentRepository := app.NewShipmentRepository()
	bloomRepository := app.NewBloomFilterRepository()
	handler := shipment.NewHandler(bloomRepository, shipmentRepository)

	shipment.AddRoutes(r, handler)

	godotenv.Load("app.env")

	go func() {
		bloomRepository.InjectFromDB()
	}()

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8000", r))
	}()

	select {}
}
