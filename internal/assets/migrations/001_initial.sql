-- +migrate Up

CREATE TABLE IF NOT EXISTS kv (
    key uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    value BYTEA NOT NULL
);

CREATE INDEX IF NOT EXISTS kv_value_idx ON kv(value);

-- +migrate Down

DROP INDEX IF EXISTS kv_value_idx;
DROP TABLE IF EXISTS kv;