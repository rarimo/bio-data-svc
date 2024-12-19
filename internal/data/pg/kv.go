package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/rarimo/bio-data-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const kvTableName = "kv"

func NewKVQ(db *pgdb.DB) data.KVQ {
	return &kvQ{
		db:       db,
		selector: sq.Select("*").From(kvTableName),
		updater:  sq.Update(kvTableName),
		deleter:  sq.Delete(kvTableName),
	}
}

type kvQ struct {
	db       *pgdb.DB
	selector sq.SelectBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
}

func (q kvQ) New() data.KVQ {
	return NewKVQ(q.db.Clone())
}

func (q kvQ) Get() (*data.KV, error) {
	var result data.KV
	err := q.db.Get(&result, q.selector)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q kvQ) Insert(data data.KV) error {
	return q.db.Exec(sq.Insert(kvTableName).SetMap(structs.Map(data)))
}

func (q kvQ) FilterByKey(key ...string) data.KVQ {
	return q.withFilters(sq.Eq{"key": key})
}

func (q kvQ) FilterByValue(value ...[]byte) data.KVQ {
	return q.withFilters(sq.Eq{"value": value})
}

func (q kvQ) withFilters(stmt interface{}) data.KVQ {
	q.selector = q.selector.Where(stmt)
	q.updater = q.updater.Where(stmt)
	q.deleter = q.deleter.Where(stmt)

	return q
}
