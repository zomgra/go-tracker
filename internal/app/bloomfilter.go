package app

import (
	"encoding/json"
	"log"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/zomgra/tracker/internal/domain"
	"github.com/zomgra/tracker/pkg/db/postgres"
)

type BloomFilterService struct {
	repository BloomFilterRepository
}
type BloomFilterRepository struct {
	OnInjecting bool
	filter      *bloom.BloomFilter
}

func NewBloomFilterRepository() *BloomFilterRepository {
	return &BloomFilterRepository{true, bloom.NewWithEstimates(100000, 0.01)}
}
func (b *BloomFilterRepository) OnLoad() bool {
	return b.OnInjecting
}

func (b *BloomFilterRepository) AddShipment(s domain.Shipment) {
	newBarcodeBytes, _ := json.Marshal(s.Barcode)
	b.filter.Add([]byte(newBarcodeBytes))
	log.Println("Use bloomfilter")

	postgres.InsertShipment(s.Barcode)
}

func (b *BloomFilterRepository) CheckShipment(barcode string) bool {
	shipmentBarcodeBytes, _ := json.Marshal(barcode)

	existInBloom := b.filter.Test(shipmentBarcodeBytes)
	if !existInBloom {
		return false
	}

	log.Println("Use bloomfilter")

	ok, err := postgres.ExistShipment(barcode)
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
	err := postgres.InjectDataTo(r.filter)
	if err != nil {
		log.Fatal(err)
	}
	r.OnInjecting = false
}
