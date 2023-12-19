package web

import (
	"github.com/gorilla/mux"
	"github.com/zomgra/tracker/internal/shipment"
)

func NewRoutes(sh *shipment.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(errorHandlerMiddleware, logMiddleware)

	r.Handle("/api/shipment", checkQuantity(sh.Create)).Methods("POST")
	r.HandleFunc("/api/shipment/{barcode}", sh.Check).Methods("GET")

	// Prometheus
	//r.Handle("/metrics", promhttp.Handler())

	return r
}
