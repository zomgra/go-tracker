package http

import (
	"github.com/gorilla/mux"
)

func AddShipmentRoutes(r *mux.Router, h *ShipmentHandler) *mux.Router {

	r.Use(errorHandlerMiddleware, logMiddleware)

	r.Handle("/api/shipment", checkQuantity(h.CreateShipments)).Methods("POST")
	r.HandleFunc("/api/shipment/{barcode}", h.CheckShipments).Methods("GET")

	// Prometheus
	//r.Handle("/metrics", promhttp.Handler())

	return r
}
