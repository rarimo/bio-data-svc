package data

type KVQ interface {
	New() KVQ

	Get() (*KV, error)
	Insert(data KV) error

	FilterByKey(key ...string) KVQ
	FilterByValue(value ...[]byte) KVQ
}

type KV struct {
	Key   string `db:"id" structs:"-"`
	Value []byte `db:"value" structs:"value"`
}
