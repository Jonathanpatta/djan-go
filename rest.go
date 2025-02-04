package djan_go

type DataModel[T any] interface {
	Get(id string) (T, error)
	Post(data T) (T, error)
	Put(data T) (T, error)
	Delete(id string) (T, error)
	List() ([]T, error)
}
