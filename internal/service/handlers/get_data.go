package handlers

import (
	"net/http"

	"github.com/rarimo/bio-data-svc/internal/data"
	"github.com/rarimo/bio-data-svc/internal/data/pg"
	"github.com/rarimo/bio-data-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDataRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	kvQuery := KVQ(r)

	if req.Key != nil {
		kvQuery = kvQuery.FilterByKey(*req.Key)
	}

	if req.Value != nil {
		kvQuery = kvQuery.
			FilterByBase64ValueLength(*req.Value).
			OrderBy(pg.HammingDistanceBase64(*req.Value), data.OrderDesc)
	}

	kv, err := kvQuery.Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get key-value")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if kv == nil {
		Log(r).Error("no key-value row found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, newKVResponse(*kv))
}
