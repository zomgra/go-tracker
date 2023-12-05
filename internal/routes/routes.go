package routes

import (
	"github.com/gorilla/mux"
	"github.com/zomgra/bitbucket/internal/shipment"
)

func CreateRoute() *mux.Router {
	r := mux.NewRouter()

	//Add Shipment route
	r.HandleFunc("/api/shipment/{id}", shipment.CheckShipments).Methods("GET")
	r.HandleFunc("/api/shipment", shipment.CreateShipments).Methods("POST")

	return r
}
