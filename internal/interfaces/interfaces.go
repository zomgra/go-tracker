package interfaces

type Repository[T any] interface {
	LoadEnding()
	Add(s T)
	Check(id string) bool
	InjectFromDB(ec chan error)
}
