package http

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zomgra/tracker/internal/domain"
	"github.com/zomgra/tracker/internal/interfaces"
)

type ShipmentHandler struct {
	bloomfilter     interfaces.Repository
	shipmentservice interfaces.Repository
}

func NewHandler(bloomfilter interfaces.Repository, shipmentservice interfaces.Repository) *ShipmentHandler {
	return &ShipmentHandler{bloomfilter: bloomfilter, shipmentservice: shipmentservice}
}

func getRepository(handler *ShipmentHandler) interfaces.Repository {
	if handler.bloomfilter.OnLoad() {
		return handler.shipmentservice
	} else {
		return handler.bloomfilter
	}
}

func (h *ShipmentHandler) CheckShipments(w http.ResponseWriter, r *http.Request) {
	log.Println(h.bloomfilter.OnLoad())

	params := mux.Vars(r)
	barcode := params["barcode"]
	ok := getRepository(h).CheckShipment(barcode)

	if ok {
		returnJson(w, ok, 200)
	} else {
		returnJson(w, "not found shipment", 404)
	}
}

func (s *ShipmentHandler) CreateShipments(w http.ResponseWriter, r *http.Request) {
	quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		http.Error(w, "Bad quantity params", http.StatusBadRequest)
	}
	shipments := make([]domain.Shipment, 0)
	for i := 0; i < quantity; i++ {
		ship := domain.Shipment{}
		ship.GenerateShipment()
		repo := getRepository(s)
		repo.AddShipment(ship)
		shipments = append(shipments, ship)
		log.Println(i)
	}
	returnJson(w, shipments, 201)
}

func returnJson(w http.ResponseWriter, v interface{}, status int) {
	if v == nil {
		w.WriteHeader(404)
		return
	}
	t := reflect.ValueOf(v)

	if t.Kind() == reflect.Slice {
		if t.Len() == 0 {
			w.WriteHeader(404)
			return
		}
	}
	//TODO: more check
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
