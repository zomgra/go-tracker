package bloomfilter

import (
	"github.com/bits-and-blooms/bloom"
)

func NewBloomFilterHelper() *BloomFilterHelper {
	bloomfilter := bloom.NewWithEstimates(1000000, 0.01)
	return &BloomFilterHelper{bloomfilter}
}

type BloomFilterHelper struct {
	filter *bloom.BloomFilter
}

func (b *BloomFilterHelper) Check(identifier []byte) bool {

	existInBloom := b.filter.Test(identifier)

	return existInBloom
}
func (b *BloomFilterHelper) Add(identifier []byte) {
	b.filter.Add(identifier)
}
