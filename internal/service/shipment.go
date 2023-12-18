package service

import (
	"encoding/json"
	"log"

	"github.com/zomgra/tracker/internal/domain"
	"github.com/zomgra/tracker/internal/interfaces"
	"github.com/zomgra/tracker/pkg/bloomfilter"
	"github.com/zomgra/tracker/pkg/db"
)

type ShipmentRepository struct {
	bloomfilterIsReady bool
	dbClient           db.Client
	bloomHelper        *bloomfilter.BloomFilterHelper
}

func NewShipmentRepository(dbClient db.Client) interfaces.Repository[domain.Shipment] {
	bloomHelper := bloomfilter.NewBloomFilterHelper()

	return &ShipmentRepository{bloomHelper: bloomHelper, dbClient: dbClient, bloomfilterIsReady: false}
}

func (r *ShipmentRepository) LoadEnding() {
	r.bloomfilterIsReady = true
}

func (r *ShipmentRepository) Add(s domain.Shipment) {
	err := r.dbClient.Insert(s.Barcode)

	if err != nil {
		log.Panic("problem with adding", err)
	}
}

func (r *ShipmentRepository) Check(id string) bool {

	barcodeByte, _ := json.Marshal(id)
	ok := r.bloomHelper.Check(barcodeByte)

	if !ok {
		return false
	}

	ok, err := r.dbClient.Exist(id)

	if err != nil {
		log.Panic("problem with checking: ", err)
	}
	log.Println("Use shipment ")

	return ok
}
func (r *ShipmentRepository) InjectFromDB(ec chan error) {
	r.bloomfilterIsReady = false
	err := r.dbClient.InjectDataTo(r.bloomHelper)
	if err != nil {
		ec <- err
		return
	}
	r.bloomfilterIsReady = true
}
