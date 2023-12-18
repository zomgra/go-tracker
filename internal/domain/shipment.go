package domain

import "github.com/zomgra/tracker/pkg/barcode"

type Shipment struct {
	Barcode string
}

func (shipment *Shipment) GenerateShipment() {
	code := barcode.GenerateBarCode(2, 2)
	shipment.Barcode = code
}
