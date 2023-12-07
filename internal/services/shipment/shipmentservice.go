package shipmentservice

import (
	"log"

	"github.com/zomgra/bitbucket/internal/models"
	"github.com/zomgra/bitbucket/pkg/db"
)

type ShipmentService struct {
	repository ShipmentRepository
}
type ShipmentRepository struct {
}

var Repository ShipmentRepository

func (r ShipmentRepository) AddShipment(s models.Shipment) {
	err := db.InsertShipment(s.Barcode)
	log.Println("Use shipment ")
	if err != nil {
		log.Panic("problem with adding", err)
	}
}

func (r ShipmentRepository) CheckShipment(barcode string) bool {
	ok, err := db.SelectShipment(barcode)
	if err != nil {
		log.Panic("problem with checking: ", err)
	}
	log.Println("Use shipment ")

	return ok
}
func (r ShipmentRepository) InjectFromDB() {
	log.Fatal("in shipment repository should not use injecting from db")
}
