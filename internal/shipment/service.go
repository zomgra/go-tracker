package shipment

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/zomgra/tracker/pkg/bloomfilter"
	"github.com/zomgra/tracker/pkg/db"
)

type Service struct {
	shipmentRepository RepositoryInterface[Shipment]
	bloomHelper        bloomfilter.BloomHelper
	bloomfilterIsReady bool

	dbClient db.Client
}

func NewService(shipmentRepository RepositoryInterface[Shipment]) *Service {
	return &Service{shipmentRepository: shipmentRepository}
}

func (s *Service) AddShipment(ship Shipment) error {
	barcodeByte, err := json.Marshal(ship.Barcode)

	if err != nil {
		return err
	}

	s.bloomHelper.Add(barcodeByte)

	return nil
}

func (s *Service) CheckShipment(barcode string) (bool, error) {

	barcodeByte, err := json.Marshal(barcode)

	if err != nil {
		return false, err
	}

	if s.bloomfilterIsReady {
		log.Println("Using bloomfilter")

		ok := s.bloomHelper.Check(barcodeByte)
		if !ok {
			return false, nil
		}
	}

	ok, err := s.shipmentRepository.Check(barcode)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (s *Service) InjectFromDB(ec chan error) {
	s.bloomfilterIsReady = false
	s.shipmentRepository.InjectFromDB(ec, s.bloomHelper)

	s.bloomfilterIsReady = true
}
