package data

type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

type KVQ interface {
	New() KVQ

	Get() (*KV, error)
	Insert(data KV) error

	FilterByKey(key string) KVQ
	FilterByValue(value []byte) KVQ
	FilterByBase64ValueLength(value string) KVQ

	OrderBy(expr interface{}, order OrderType) KVQ
}

type KV struct {
	Key   string `db:"key" structs:"-"`
	Value []byte `db:"value" structs:"value"`
}
