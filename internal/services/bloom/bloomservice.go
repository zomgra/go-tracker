package bloomfilter

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/zomgra/bitbucket/internal/models"
	shipmentservice "github.com/zomgra/bitbucket/internal/services/shipment"
)

var filter *bloom.BloomFilter = bloom.NewWithEstimates(1000000, 0.01)

func AddShipment(s models.Shipment) {
	newBarcodeBytes, _ := json.Marshal(s.Barcode)
	filter.Add([]byte(newBarcodeBytes))
	shipmentservice.AddShipment(s.Barcode)
}
func CheckShipment(barcode string) bool {
	shipmentBarcodeBytes, _ := json.Marshal(barcode)
	existInBloom := filter.Test(shipmentBarcodeBytes)
	if !existInBloom {
		log.Panic(errors.New("not exist in bloom filter"))
	}
	ok := shipmentservice.CheckShipment(barcode)
	if ok {
		return true
	}
	log.Panic(errors.New("not found shipment"))
	return false
}
