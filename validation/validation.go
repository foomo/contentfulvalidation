package validation

import (
	"github.com/foomo/contentfulvalidation/errors"
	"github.com/foomo/contentfulvalidation/validator"
)

type Validation interface {
	ValidationResult(modelType validator.ModelType, modelID validator.ModelID) (validationResult *validator.ValidationResult, validationError *errors.ValidationError)
	ValidationResults(modelType validator.ModelType) (validationResults map[validator.ModelID]*validator.ValidationResult, validationError *errors.ValidationError)
	ValidateEntity(modelType validator.ModelType, modelID validator.ModelID, commit bool) (validationResult *validator.ValidationResult, validationError *errors.ValidationError)
	ListModelTypes() (availableModelTypes []*validator.ModelTypeInfo)
}
