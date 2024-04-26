package webhook

import (
	"net/http"

	"github.com/foomo/contentfulvalidation/validator"
)

type Webhook interface {
	UpdateCache(w http.ResponseWriter, r *http.Request, sysType validator.SysType, modelType validator.ModelType, modelID validator.ModelID)
}
