package interfaces

type Inter interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Check(key string) (bool, error)
}
