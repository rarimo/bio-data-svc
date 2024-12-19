package handlers

import (
	"context"
	"net/http"

	"github.com/rarimo/bio-data-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	kvCtxKey  ctxKey = iota
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxKVQ(stateQ data.KVQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, kvCtxKey, stateQ)
	}
}

func KVQ(r *http.Request) data.KVQ {
	return r.Context().Value(kvCtxKey).(data.KVQ)
}
