package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/google/uuid"
	"github.com/rarimo/bio-data-svc/internal/data"
	"github.com/rarimo/bio-data-svc/internal/service/requests"
	"github.com/rarimo/bio-data-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AddData(w http.ResponseWriter, r *http.Request) {
	req, err := requests.AddData(r)
	if err != nil {
		Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	value, err := base64.StdEncoding.DecodeString(req.Data.Attributes.Value)
	if err != nil {
		Log(r).WithError(err).WithField("value", req.Data.Attributes.Value).Error("failed to decode base64 string")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	kv := data.KV{
		Key:   uuid.NewString(),
		Value: value,
	}

	if err = KVQ(r).Insert(kv); err != nil {
		Log(r).WithError(err).WithField("kv", kv).Error("failed to insert new key-value")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, newKVResponse(kv))
}

func newKVResponse(kv data.KV) resources.ValueResponse {
	return resources.ValueResponse{
		Data: resources.Value{
			Key: resources.Key{
				ID:   kv.Key,
				Type: resources.VALUE,
			},
			Attributes: resources.ValueAttributes{
				Value: base64.StdEncoding.EncodeToString(kv.Value),
				Key:   kv.Key,
			},
		},
	}
}
