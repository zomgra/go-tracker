package shipment

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zomgra/bitbucket/internal/models"
	bloomfilter "github.com/zomgra/bitbucket/internal/services/bloom"
)

func CheckShipments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	barcode := params["barcode"]

	ok := bloomfilter.CheckShipment(barcode)

	if ok {
		returnJson(w, ok, 200)
	} else {
		returnJson(w, "not found shipment", 404)
	}
}
func CreateShipments(w http.ResponseWriter, r *http.Request) {
	quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		http.Error(w, "Bad quantity params", http.StatusBadRequest)
	}
	log.Println(quantity)
	var shipments []models.Shipment
	for i := 0; i < quantity; i++ {

		s := models.Shipment{}
		s.GenerateShipment()
		bloomfilter.AddShipment(s)
		shipments = append(shipments, s)
	}

	returnJson(w, shipments, 201)
}

func returnJson(w http.ResponseWriter, v interface{}, status int) {
	if v == nil {
		// No Content
	}
	t := reflect.ValueOf(v)

	if t.Kind() == reflect.Slice {
		if t.Len() == 0 {
			//TODO: Return: No Content
		}
	}
	//TODO: more check
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
