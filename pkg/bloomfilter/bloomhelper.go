package bloomfilter

import (
	"encoding/json"

	"github.com/bits-and-blooms/bloom"
)

func NewBloomFilterHelper() *Helper {
	bloomfilter := bloom.NewWithEstimates(1000000, 0.01)
	return &Helper{bloomfilter}
}

type Helper struct {
	filter *bloom.BloomFilter
}

func (b *Helper) Check(identifier []byte) bool {

	existInBloom := b.filter.Test(identifier)

	return existInBloom
}
func (b *Helper) Add(identifier []byte) {
	b.filter.Add(identifier)
}

func (b *Helper) Inject(f func(chan any) error) error {
	c := make(chan any)
	errChan := make(chan error)
	go func() {
		err := f(c)
		errChan <- err
	}()
	for value := range c {
		barcodeByte, err := json.Marshal(value)
		if err != nil {
			return err
		}
		b.Add(barcodeByte)
	}
	return <-errChan
}
