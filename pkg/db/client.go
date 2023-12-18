package db

import (
	"github.com/zomgra/tracker/pkg/bloomfilter"
)

type Client interface {
	Insert(string) error
	Exist(barcode string) (bool, error)
	InjectDataTo(filter *bloomfilter.BloomFilterHelper) error
}
