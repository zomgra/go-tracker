package bloomfilter

type BloomHelper interface {
	Check(identifier []byte) bool
	Add(identifier []byte)
	Inject(f func(chan any) error) error
}
