package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/is"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/zk-biometrics-svc/resources"
)

func AddData(r *http.Request) (req resources.AddValueRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, val.Errors{"body": err}
	}

	return req, val.Errors{
		"data/type":             val.Validate(req.Data.Type, val.Required, val.In(resources.VALUE)),
		"data/attributes/value": val.Validate(req.Data.Attributes.Value, val.Required, is.Base64),
	}.Filter()
}
