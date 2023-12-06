package bloomfilter

import (
	"encoding/json"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/zomgra/bitbucket/internal/models"
)

var filter *bloom.BloomFilter = bloom.NewWithEstimates(1000000, 0.01)

func GetFilter() *bloom.BloomFilter {
	return filter
}
func AddShipment(s models.Shipment) {
	newBarcodeBytes, _ := json.Marshal(s.Barcode)
	filter.Add([]byte(newBarcodeBytes))
}
func CheckShipment(barcode string) bool {
	shipmentBarcodeBytes, _ := json.Marshal(barcode)
	filter.Test(shipmentBarcodeBytes)
}
