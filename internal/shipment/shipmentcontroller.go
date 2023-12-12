package shipment

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zomgra/bitbucket/internal/interfaces"
	"github.com/zomgra/bitbucket/internal/models"

	bloomfilter "github.com/zomgra/bitbucket/internal/services/bloom"
	shipmentservice "github.com/zomgra/bitbucket/internal/services/shipment"
)

func getRepository() interfaces.Repository {

	if bloomfilter.Repository.OnLoad {
		return shipmentservice.Repository
	}
	return bloomfilter.Repository
}

func CheckShipments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	barcode := params["barcode"]
	ok := getRepository().CheckShipment(barcode)

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
	shipments := make([]models.Shipment, 0)
	for i := 0; i < quantity; i++ {
		s := models.Shipment{}
		s.GenerateShipment()
		getRepository().AddShipment(s)
		shipments = append(shipments, s)
		log.Println(i)
	}
	returnJson(w, shipments, 201)
	//shipments = nil
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
