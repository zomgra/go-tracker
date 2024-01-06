package bloomfilter

import (
	"encoding/json"

	"github.com/bits-and-blooms/bloom"
	"github.com/zomgra/tracker/pkg/config"
)

func NewBloomFilterHelper(c *config.BloomFilterConfig) *Helper {
	bloomfilter := bloom.NewWithEstimates(c.BloomLimit, 0.01)
	return &Helper{bloomfilter}
}

// make interface
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
