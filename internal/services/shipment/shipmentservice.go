package shipmentservice

import (
	"log"

	"github.com/zomgra/bitbucket/pkg/db"
)

func AddShipment(barcode string) {
	err := db.InsertShipment(barcode)
	if err != nil {
		log.Panic("problem with adding", err)
	}
}

func CheckShipment(barcode string) bool {
	ok, err := db.SelectShipment(barcode)
	if err != nil {
		log.Panic("problem with checking: ", err)
	}
	return ok
}
