package shipment

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router, h *ShipmentHandler) *mux.Router {

	r.Use(errorHandlerMiddleware, logMiddleware)
	//Add Shipment route
	r.Handle("/api/shipment", checkQuantity(h.CreateShipments)).Methods("POST")
	r.HandleFunc("/api/shipment/{barcode}", h.CheckShipments).Methods("GET")

	// Prometheus
	//r.Handle("/metrics", promhttp.Handler())

	return r
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s - %s (%s) Time beetween excecuting: %v \n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func checkQuantity(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid quantity: %s", err), http.StatusBadRequest)
			return
		}
		if quantity < 0 {
			http.Error(w, "Quantity must be greater than or equal to 0", http.StatusBadRequest)
			return
		}

		if f != nil {
			f(w, r)
		} else {
			log.Println("checkQuantity middleware: Handler function is nil")
		}
	})
}

func errorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if w != nil {
					http.Error(w, fmt.Sprintf("panic on the server: %s", err), http.StatusInternalServerError)
				} else {
					log.Printf("Panic on the server: %s", err)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}