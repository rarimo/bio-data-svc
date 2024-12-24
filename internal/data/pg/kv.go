package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/rarimo/zk-biometrics-svc/internal/data"
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

func (q kvQ) Delete() error {
	return q.db.Exec(q.deleter)
}

func (q kvQ) FilterByKey(key string) data.KVQ {
	return q.withFilters(sq.Eq{"key": key})
}

func (q kvQ) FilterByValue(value []byte) data.KVQ {
	return q.withFilters(sq.Eq{"value": value})
}

func (q kvQ) FilterByBase64ValueLength(value string) data.KVQ {
	return q.withFilters(sq.Expr("LENGTH(value) = LENGTH(DECODE(?, 'base64'))", value))
}

func (q kvQ) OrderBy(expr sq.Sqlizer, order data.OrderType) data.KVQ {
	q.selector = q.selector.OrderByClause(expr, string(order))

	return q
}

func (q kvQ) withFilters(stmt interface{}) data.KVQ {
	q.selector = q.selector.Where(stmt)
	q.updater = q.updater.Where(stmt)
	q.deleter = q.deleter.Where(stmt)

	return q
}

func HammingDistanceBase64(value string) sq.Sqlizer {
	return sq.Expr("hamming_distance(value, DECODE(?, 'base64'))", value)
}
