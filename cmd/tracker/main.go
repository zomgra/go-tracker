package main

import (
	"log"
	serverhttp "net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/zomgra/tracker/internal/app"
	"github.com/zomgra/tracker/internal/http"
)

func main() {
	r := mux.NewRouter()
	shipmentRepository := app.NewShipmentRepository()
	bloomRepository := app.NewBloomFilterRepository()
	handler := http.NewHandler(bloomRepository, shipmentRepository)

	http.AddShipmentRoutes(r, handler)

	godotenv.Load("app.env")

	go func() {
		bloomRepository.InjectFromDB()
	}()

	go func() {
		log.Fatal(serverhttp.ListenAndServe("localhost:8000", r))
	}()

	select {}
}
