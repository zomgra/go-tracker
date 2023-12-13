package bloomfilter

import (
	"encoding/json"
	"log"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/zomgra/tracker/internal/models"
	"github.com/zomgra/tracker/pkg/db"
)

type BloomFilterService struct {
	repository BloomFilterRepository
}
type BloomFilterRepository struct {
	OnInjecting bool
	filter      *bloom.BloomFilter
}

func NewRepository() *BloomFilterRepository {
	return &BloomFilterRepository{true, bloom.NewWithEstimates(100000, 0.01)}
}
func (b *BloomFilterRepository) OnLoad() bool {
	return b.OnInjecting
}

func (b *BloomFilterRepository) AddShipment(s models.Shipment) {
	newBarcodeBytes, _ := json.Marshal(s.Barcode)
	b.filter.Add([]byte(newBarcodeBytes))
	log.Println("Use bloomfilter")

	db.InsertShipment(s.Barcode)
}

func (b *BloomFilterRepository) CheckShipment(barcode string) bool {
	shipmentBarcodeBytes, _ := json.Marshal(barcode)

	existInBloom := b.filter.Test(shipmentBarcodeBytes)
	if !existInBloom {
		return false
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
	r.OnInjecting = true
	err := db.InjectDataTo(r.filter)
	if err != nil {
		log.Fatal(err)
	}
	r.OnInjecting = false
}
