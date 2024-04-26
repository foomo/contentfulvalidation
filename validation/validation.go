package validation

import (
	"net/http"

	"github.com/foomo/contentfulvalidation/errors"
	"github.com/foomo/contentfulvalidation/validator"
)

type Validation interface {
	ValidationResult(w http.ResponseWriter, r *http.Request, modelType validator.ModelType, modelID validator.ModelID) (validationResult *validator.ValidationResult, validationError *errors.ValidationError)
	ValidationResults(w http.ResponseWriter, r *http.Request, modelType validator.ModelType) (validationResults map[validator.ModelID]*validator.ValidationResult, validationError *errors.ValidationError)
	ValidateEntity(w http.ResponseWriter, r *http.Request, modelType validator.ModelType, modelID validator.ModelID, commit bool) (validationResult *validator.ValidationResult, validationError *errors.ValidationError)
	ListModelTypes(w http.ResponseWriter, r *http.Request) (availableModelTypes []*validator.ModelTypeInfo)
}
