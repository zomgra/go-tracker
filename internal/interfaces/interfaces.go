package interfaces

type Repository[T any] interface {
	LoadEnding()
	Add(s T) error
	Check(id string) (bool, error)
	InjectFromDB(ec chan error)
}
