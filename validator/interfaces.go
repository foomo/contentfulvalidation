package validator

import (
	"context"

	"github.com/foomo/contentfulvalidation/errors"
)

type ModelValidator interface {
	Validate(ctx context.Context, modelID ModelID) (*ValidationResult, *errors.ValidationError)
	ValidateAll(ctx context.Context) (map[ModelID]*ValidationResult, error)
}

type ValidatorProvider interface {
	GetValidators() map[ModelType]ModelValidator
}
