package routes

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/zomgra/bitbucket/internal/shipment"
)

func CreateRoute() *mux.Router {
	r := mux.NewRouter()

	r.Use(logMiddlware)

	//Add Shipment route
	r.HandleFunc("/api/shipment/{id}", shipment.CheckShipments).Methods("GET")
	r.Handle("/api/shipment", checkQuantity(shipment.CreateShipments)).Methods("POST")

	return r
}

func logMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s - %s (%s) Time beetween excecuting: %v \n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start).Milliseconds())
	})
}

func checkQuantity(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
		if err != nil {
			log.Panic(err)
		}
		if quantity < 0 {
			log.Fatal("quantity lower 0")
		}
		f(w, r)
	})
}
