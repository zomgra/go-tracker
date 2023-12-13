package bloomfilter

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/zomgra/bitbucket/internal/models"
	"github.com/zomgra/bitbucket/pkg/db"
)

// TODO : Redesign services on interfaces, for using once if other unhealty

type BloomFilterService struct {
	repository BloomFilterRepository
}
type BloomFilterRepository struct {
	OnLoad bool
}

var Repository BloomFilterRepository = BloomFilterRepository{OnLoad: true}

var filter *bloom.BloomFilter = bloom.NewWithEstimates(1000000, 0.01)

func (b BloomFilterRepository) AddShipment(s models.Shipment) {
	newBarcodeBytes, _ := json.Marshal(s.Barcode)
	filter.Add([]byte(newBarcodeBytes))
	log.Println("Use bloomfilter")

	// TODO : Change it so that it is not necessary to use ShipmentRepository

	db.InsertShipment(s.Barcode)
}
func (b BloomFilterRepository) CheckShipment(barcode string) bool {
	shipmentBarcodeBytes, _ := json.Marshal(barcode)

	existInBloom := filter.Test(shipmentBarcodeBytes)
	if !existInBloom {
		log.Panic(errors.New("not exist in bloom filter"))
	}

	// TODO : Change it so that it is not necessary to use ShipmentRepository
	log.Println("Use bloomfilter")

	ok, err := db.ExistShipment(barcode)
	if err != nil {
		return false
	}
	if ok {
		return true
	}
	return false
}
func (r BloomFilterRepository) InjectFromDB() {
	Repository.OnLoad = true
	db.InjectDataTo(filter)
	Repository.OnLoad = false
}
