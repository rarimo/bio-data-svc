package data

import sq "github.com/Masterminds/squirrel"

type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

type KVQ interface {
	New() KVQ

	Get() (*KV, error)
	Insert(data KV) error
	Delete() error

	FilterByKey(key string) KVQ
	FilterByValue(value []byte) KVQ
	FilterByBase64ValueLength(value string) KVQ

	OrderBy(expr sq.Sqlizer, order OrderType) KVQ
}

type KV struct {
	Key   string `db:"key" structs:"-"`
	Value []byte `db:"value" structs:"value"`
}
