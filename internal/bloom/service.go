package bloomfilter

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/zomgra/tracker/internal/models"
	"github.com/zomgra/tracker/pkg/db"
)

// TODO : Redesign services on interfaces, for using once if other unhealty

type BloomFilterService struct {
	repository BloomFilterRepository
}
type BloomFilterRepository struct {
	OnLoad bool
	filter *bloom.BloomFilter
}

func NewRepository() *BloomFilterRepository {
	return &BloomFilterRepository{true, bloom.NewWithEstimates(100000, 0.01)}
}

func (b *BloomFilterRepository) AddShipment(s models.Shipment) {
	newBarcodeBytes, _ := json.Marshal(s.Barcode)
	b.filter.Add([]byte(newBarcodeBytes))
	log.Println("Use bloomfilter")

	// TODO : Change it so that it is not necessary to use ShipmentRepository

	db.InsertShipment(s.Barcode)
}
func (b *BloomFilterRepository) CheckShipment(barcode string) bool {
	shipmentBarcodeBytes, _ := json.Marshal(barcode)

	existInBloom := b.filter.Test(shipmentBarcodeBytes)
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
func (r *BloomFilterRepository) InjectFromDB() {
	r.OnLoad = true
	db.InjectDataTo(r.filter)
	r.OnLoad = false
}
