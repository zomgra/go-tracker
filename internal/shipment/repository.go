package shipment

import (
	"encoding/json"
	"log"

	"github.com/zomgra/tracker/internal/interfaces"
	"github.com/zomgra/tracker/pkg/bloomfilter"
	"github.com/zomgra/tracker/pkg/db"
)

type ShipmentRepository struct {
	bloomfilterIsReady bool
	dbClient           db.Client
	bloomHelper        *bloomfilter.Helper
}

func NewRepository(dbClient db.Client) interfaces.Repository[Shipment] {
	bloomHelper := bloomfilter.NewBloomFilterHelper()

	return &ShipmentRepository{bloomHelper: bloomHelper, dbClient: dbClient, bloomfilterIsReady: false}
}

func (r *ShipmentRepository) LoadEnding() {
	r.bloomfilterIsReady = true
}

func (r *ShipmentRepository) Add(s Shipment) error {

	barcodeByte, err := json.Marshal(s.Barcode)

	if err != nil {
		return err
	}

	r.bloomHelper.Add(barcodeByte)

	err = r.dbClient.Insert(s.Barcode)

	if err != nil {
		return err
	}
	return nil
}

func (r *ShipmentRepository) Check(id string) (bool, error) {

	barcodeByte, _ := json.Marshal(id)
	if r.bloomfilterIsReady {
		log.Println("Using bloomfilter")

		ok := r.bloomHelper.Check(barcodeByte)
		if !ok {
			return false, nil
		}
	}

	ok, err := r.dbClient.Exists(id)

	if err != nil {
		return false, err
	}
	log.Println("Use shipment ")

	return ok, nil
}
func (r *ShipmentRepository) InjectFromDB(ec chan error) {
	r.bloomfilterIsReady = false
	err := r.bloomHelper.Inject(r.dbClient.InjectDataTo)
	if err != nil {
		ec <- err
		return
	}
	r.bloomfilterIsReady = true
}
