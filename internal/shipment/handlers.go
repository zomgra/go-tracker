package shipment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	barcode := params["barcode"]
	ok, err := h.service.CheckShipment(barcode)

	if ok {
		returnJson(w, ok, 200)
	} else {
		returnJson(w, &Error{"not found shipment: ", err}, 404)
	}
}

func (s *Handler) Create(w http.ResponseWriter, r *http.Request) {
	quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		http.Error(w, "Bad quantity params", http.StatusBadRequest)
	}
	shipments := make([]Shipment, 0)
	for i := 0; i < quantity; i++ {
		ship := Shipment{}
		ship.GenerateShipment()

		err := s.service.AddShipment(ship)
		if err != nil {
			returnJson(w, fmt.Sprintf("error with creating shipment: %v ", err), 500)
		}
		shipments = append(shipments, ship)
	}
	returnJson(w, shipments, 201)
}

func returnJson(w http.ResponseWriter, v any, status int) {
	if v == nil {
		w.WriteHeader(404)
		return
	}
	t := reflect.ValueOf(v)
	err, ok := v.(*Error)
	if ok {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
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
