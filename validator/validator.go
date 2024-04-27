package validator

import (
	"context"
	"strings"
	"time"

	"github.com/foomo/contentfulvalidation/constants"
	"github.com/foomo/contentfulvalidation/errors"
	"github.com/foomo/contentfulvalidation/utils"
	"github.com/foomo/contentserver/client"
	keellog "github.com/foomo/keel/log"
	"go.uber.org/zap"
)

type Validator struct {
	l          *zap.Logger
	csClient   *client.Client
	Validators map[ModelType]ModelValidator
	Cache      *Cache
}

func NewValidator(
	l *zap.Logger,
	csClient *client.Client,
	validatorProvider ValidatorProvider,
) (*Validator, error) {
	logger := l.With(zap.String("routine", "contentfulvalidation-validator"))
	cache, err := NewCache(l)
	if err != nil {
		return nil, err
	}

	return &Validator{
		l:          logger,
		csClient:   csClient,
		Validators: validatorProvider.GetValidators(),
		Cache:      cache,
	}, nil
}

func (v *Validator) Get(modelType ModelType, modelID ModelID) (
	*ValidationResult,
	*errors.ValidationError,
) {
	result, ok := v.Cache.Get(modelType, modelID)
	if !ok {
		return nil, nil
	}
	return result, nil
}

func (v *Validator) Validate(ctx context.Context, modelType ModelType, modelID ModelID) (
	*ValidationResult,
	*errors.ValidationError,
) {
	// select validator
	validator, err := v.getValidatorByType(modelType)
	if err != nil {
		return nil, err
	}
	result, err := validator.Validate(ctx, modelID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *Validator) List(modelType ModelType) (
	map[ModelID]*ValidationResult,
	error,
) {
	v.l.Debug("getting all validation results for model type", keellog.FValue(modelType))
	results := v.Cache.GetForType(modelType)
	v.l.Debug("got validation results", keellog.FValue(modelType), keellog.FValue(len(results)))
	return results, nil
}

func (v *Validator) MapListToCSV(modelType ModelType) (
	string,
	error,
) {
	r, err := v.List(modelType)
	if err != nil {
		return "", err
	}

	csvString := ""
	rows := []string{"ID;Link;Element;Internal Title;Last Updated;Status;Details"}

	for i := range r {
		v, err := v.Get(modelType, i)
		if err != nil {
			return "", err
		}

		if v.Health != constants.HealthOk {
			date, err := utils.ConvertTimeFormat(v.LastUpdatedDate, time.RFC3339, constants.DateFormat)
			if err != nil {
				return "", err
			}
			idString := string(v.ID)
			hyperLink := `=HYPERLINK("https://app.contentful.com/spaces/qfsyzz7ytbcy/entries/` + idString + `")`
			details := []string{}
			for _, m := range v.Messages {
				details = append(details, m.Message)
			}
			cells := []string{
				idString,
				hyperLink,
				v.Title,
				v.InternalTitle,
				date,
				string(v.Health),
				strings.Join(details, " "),
			}
			rows = append(rows, strings.Join(cells, ";"))
			csvString = strings.Join(rows, "\n")
		}
	}

	return csvString, nil
}

func (v *Validator) ListModelTypes() []ModelType {
	availableTypes := []ModelType{}
	for modelType := range v.Cache.GetPool() {
		availableTypes = append(availableTypes, modelType)
	}
	return availableTypes
}

func (v *Validator) ValidateAll(ctx context.Context) error {
	v.l.Debug("running validation on all model types")
	for modelType, modelValidator := range v.Validators {
		results, err := modelValidator.ValidateAll(ctx)
		if err != nil {
			keellog.WithError(v.l, err).Error("error on running validation", keellog.FValue(modelType))
			continue
		}
		v.l.Debug("successful validation run", keellog.FValue(modelType), keellog.FValue(len(results)))
		// set the whole result to the cache
		v.Cache.SetForType(modelType, results)
	}
	return nil
}

func (v *Validator) Update(ctx context.Context) {
	v.l.Debug("received update signal")
	err := v.ValidateAll(ctx)
	if err != nil {
		keellog.WithError(v.l, err).Error("error on validate all update")
	}
}

func (v *Validator) getValidatorByType(modelType ModelType) (ModelValidator, *errors.ValidationError) {
	validator, ok := v.Validators[modelType]
	if !ok {
		return nil, errors.ErrValidationErrorUnknownType
	}
	return validator, nil
}

func InitEmptyValidationResult(modelType ModelType, modelID ModelID, title string, internalTitle string, lastUpdatedDate string) *ValidationResult {
	var messages []*ValidationResultMessage
	return &ValidationResult{
		ID:              modelID,
		Title:           title,
		InternalTitle:   internalTitle,
		LastUpdatedDate: lastUpdatedDate,
		ModelType:       modelType,
		Health:          constants.HealthOk,
		Messages:        messages,
	}
}
