package shipment

import "github.com/zomgra/tracker/pkg/bloomfilter"

type RepositoryInterface[T any] interface {
	InjectFromDB(ec chan error, h bloomfilter.BloomHelper)
	Add(s T) error
	Check(id string) (bool, error)
}
