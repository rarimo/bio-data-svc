package requests

import (
	"net/http"

	"github.com/go-ozzo/ozzo-validation/is"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval/v4"
)

type GetDataRequest struct {
	Key   *string `filter:"key"`
	Value *string `filter:"value"`
}

func NewGetDataRequest(r *http.Request) (req GetDataRequest, err error) {
	if err = urlval.Decode(r.URL.Query(), &req); err != nil {
		return req, val.Errors{"query": err}
	}

	err = val.Errors{
		"key":   val.Validate(req.Key, val.When(!val.IsEmpty(req.Key), is.UUID)),
		"value": val.Validate(req.Value, val.When(!val.IsEmpty(req.Value), is.Base64)),
	}.Filter()
	return
}
