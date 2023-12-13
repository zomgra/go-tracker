package app

import (
	"log"

	"github.com/zomgra/tracker/internal/domain"
	"github.com/zomgra/tracker/pkg/db/postgres"
)

type ShipmentService struct {
	repository ShipmentRepository
}
type ShipmentRepository struct {
}

func NewShipmentRepository() *ShipmentRepository {
	return &ShipmentRepository{}
}

func (r *ShipmentRepository) OnLoad() bool {
	return false
}

func (r *ShipmentRepository) AddShipment(s domain.Shipment) {
	err := postgres.InsertShipment(s.Barcode)
	log.Println("Use shipment ")
	if err != nil {
		log.Panic("problem with adding", err)
	}
}

func (r *ShipmentRepository) CheckShipment(barcode string) bool {
	ok, err := postgres.ExistShipment(barcode)
	if err != nil {
		log.Panic("problem with checking: ", err)
	}
	log.Println("Use shipment ")

	return ok
}
func (r *ShipmentRepository) InjectFromDB() {
	log.Fatal("in shipment repository should not use injecting from db")
}
