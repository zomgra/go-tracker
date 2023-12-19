package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zomgra/tracker/internal/domain"
	"github.com/zomgra/tracker/internal/interfaces"
)

type ShipmentHandler struct {
	shipmentRepository interfaces.Repository[domain.Shipment]
}

func NewHandler(shipmentRepository interfaces.Repository[domain.Shipment]) *ShipmentHandler {
	return &ShipmentHandler{shipmentRepository: shipmentRepository}
}

func (h *ShipmentHandler) CheckShipments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	barcode := params["barcode"]
	ok, err := h.shipmentRepository.Check(barcode)

	if ok {
		returnJson(w, ok, 200)
	} else {
		returnJson(w, fmt.Sprintf("not found shipment: %v ", err), 404)
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

		err := s.shipmentRepository.Add(ship)
		if err != nil {
			returnJson(w, fmt.Sprintf("error with creating shipment: %v ", err), 500)
		}
		shipments = append(shipments, ship)
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
