package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/rarimo/zk-biometrics-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDeleteDataRequest(r)
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
		value, err := base64.StdEncoding.DecodeString(*req.Value)
		if err != nil {
			Log(r).WithError(err).WithField("value", *req.Value).Error("failed to decode Base64 string")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		kvQuery = kvQuery.FilterByValue(value)
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

	if err = kvQuery.Delete(); err != nil {
		Log(r).WithError(err).Error("failed to delete key-value")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
