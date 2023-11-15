package validator

import "github.com/foomo/contentfulvalidation/errors"

type ModelValidator interface {
	Validate(modelID ModelID) (*ValidationResult, *errors.ValidationError)
	ValidateAll() (map[ModelID]*ValidationResult, error)
}

type ValidatorProvider interface { //nolint:revive
	GetValidators() map[ModelType]ModelValidator
}
