package utils

import (
	"encoding/json"
	catvo "github.com/bestbytes/catalogue/vo"
	"github.com/foomo/contentfulvalidation/contants"
	"github.com/foomo/contentfulvalidation/validator"
	"time"

	"github.com/foomo/contentful"
	"github.com/pkg/errors"
)

func GetAssetImage(asset *contentful.AssetNoLocale) *contentful.FileImage {
	if asset != nil && asset.Fields != nil && asset.Fields.File != nil && asset.Fields.File.Detail != nil && asset.Fields.File.Detail.Image != nil {
		return asset.Fields.File.Detail.Image
	}
	return nil
}

func GetAspectRatio(asset *contentful.AssetNoLocale) (float64, error) {
	var aspectRatio float64
	image := GetAssetImage(asset)
	if image == nil {
		return aspectRatio, errors.New("No linked image available")
	}
	if image.Width == 0 || image.Height == 0 {
		return aspectRatio, errors.New("Width or height are zero")
	}
	aspectRatio = float64(image.Width) / float64(image.Height)
	return aspectRatio, nil
}

func LoadQuery(rawQuery any) (*catvo.Query, error) {
	query := &catvo.Query{}
	errMarshal := loadInterfaceAsJSON(rawQuery, query)
	if errMarshal != nil {
		return nil, errMarshal
	}
	return query, nil
}
func loadInterfaceAsJSON(source interface{}, target interface{}) error {
	jsonBytes, errMarshal := json.Marshal(source)
	if errMarshal != nil {
		return errMarshal
	}
	return json.Unmarshal(jsonBytes, &target)
}

func ConvertTimeFormat(timeToFormat string, parseTemplate string, formatTemplate string) (string, error) {
	p, err := time.Parse(parseTemplate, timeToFormat)
	if err != nil {
		return "", err
	}
	return p.Format(formatTemplate), nil
}

func InitEmptyValidationResult(modelType validator.ModelType, modelID validator.ModelID, title string, internalTitle string, lastUpdatedDate string) *validator.ValidationResult {
	var messages []*validator.ValidationResultMessage
	return &validator.ValidationResult{
		ID:              modelID,
		Title:           title,
		InternalTitle:   internalTitle,
		LastUpdatedDate: lastUpdatedDate,
		ModelType:       modelType,
		Health:          contants.HealthOk,
		Messages:        messages,
	}
}
