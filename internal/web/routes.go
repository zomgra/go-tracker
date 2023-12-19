package web

import (
	"github.com/gorilla/mux"
	"github.com/zomgra/tracker/internal/handlers"
)

func NewRoutes(h *handlers.ShipmentHandler) *mux.Router {
	r := mux.NewRouter()

	r.Use(errorHandlerMiddleware, logMiddleware)

	r.Handle("/api/shipment", checkQuantity(h.CreateShipments)).Methods("POST")
	r.HandleFunc("/api/shipment/{barcode}", h.CheckShipments).Methods("GET")

	// Prometheus
	//r.Handle("/metrics", promhttp.Handler())

	return r
}
