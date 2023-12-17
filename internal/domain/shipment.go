package domain

import "github.com/zomgra/tracker/utils"

type Shipment struct {
	Barcode string
}

type Repository interface {
	OnLoad() bool
	AddShipment(Shipment)
	CheckShipment(barcode string) bool
	InjectFromDB()
}

func (shipment *Shipment) GenerateShipment() {
	code := utils.GenerateBarCode(2, 2)
	shipment.Barcode = code
}
