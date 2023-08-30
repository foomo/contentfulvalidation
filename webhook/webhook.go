package webhook

import (
	"github.com/foomo/contentfulvalidation/validator"
)

type Webhook interface {
	UpdateCache(sysType validator.SysType, modelType validator.ModelType, modelID validator.ModelID)
}
