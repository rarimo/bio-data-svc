package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval/v4"
)

const FilterQueryErrMsg = "Either filter[key] or filter[value] is required."

type DeleteDataRequest struct {
	Key   *string `filter:"key"`
	Value *string `filter:"value"`
}

func NewDeleteDataRequest(r *http.Request) (req GetDataRequest, err error) {
	if err = urlval.Decode(r.URL.Query(), &req); err != nil {
		return req, val.Errors{"query": err}
	}

	err = val.Errors{
		"key": val.Validate(req.Key,
			val.When(!val.IsEmpty(req.Key), is.UUID),
			val.When(val.IsEmpty(req.Value), validation.Required.Error(FilterQueryErrMsg)),
		),
		"value": val.Validate(req.Value,
			val.When(!val.IsEmpty(req.Value), is.Base64),
			val.When(val.IsEmpty(req.Key), validation.Required.Error(FilterQueryErrMsg)),
		),
	}.Filter()
	return
}
