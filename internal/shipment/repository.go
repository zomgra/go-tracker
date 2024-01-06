package shipment

import (
	log "github.com/sirupsen/logrus"

	"github.com/zomgra/tracker/pkg/bloomfilter"
	"github.com/zomgra/tracker/pkg/db"
)

// CLEAR BLOOM
type Repository struct {
	dbClient db.Client
}

func NewRepository(dbClient db.Client) RepositoryInterface[Shipment] {

	return &Repository{dbClient: dbClient}
}

func (r *Repository) Add(s Shipment) error {

	err := r.dbClient.Insert(s.Barcode)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Check(id string) (bool, error) {

	ok, err := r.dbClient.Exists(id)

	if err != nil {
		return false, err
	}
	log.Println("Use shipment ")

	return ok, nil
}
func (r *Repository) InjectFromDB(ec chan error, h bloomfilter.BloomHelper) {
	err := h.Inject(r.dbClient.InjectDataTo)
	if err != nil {
		ec <- err
		return
	}
}
