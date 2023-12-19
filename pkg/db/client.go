package db

type Client interface {
	Insert(string) error
	Exists(barcode string) (bool, error)
	InjectDataTo(chan any) error
}
