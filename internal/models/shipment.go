package models

import "github.com/zomgra/bitbucket/utils"

type Shipment struct {
	Barcode string
}

func (shipment *Shipment) GenerateShipment() {
	code := utils.GenerateBarCode(2, 2)
	shipment.Barcode = code
}
