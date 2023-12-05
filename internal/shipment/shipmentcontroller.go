package shipment

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckShipments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
}
func CreateShipments(w http.ResponseWriter, r *http.Request) {
	quantity := r.URL.Query().Get("quantity")
	log.Println(quantity)
}
