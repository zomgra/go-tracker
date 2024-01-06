package web

import (
	"github.com/gorilla/mux"
	"github.com/zomgra/tracker/internal/shipment"
)

func NewRoutes(sh *shipment.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(logMiddleware)

	r.HandleFunc("/api/shipment", sh.Create).Methods("POST")
	r.HandleFunc("/api/shipment/{barcode}", sh.Check).Methods("GET")

	return r
}
