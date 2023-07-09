package nosql

type PrimaryKey string
type Document map[string]interface{}

type Store interface {
	Exists(PrimaryKey) (bool, error)
	Set(Document) error
	Get(PrimaryKey, interface{}) error
	List() ([]PrimaryKey, error)
	Remove(PrimaryKey) error
	Clear() error
}
