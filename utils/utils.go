package utils

import (
	"encoding/json"
	"time"

	"github.com/foomo/contentful"
	"github.com/pkg/errors"
)

type imageSize struct {
	Width  int
	Height int
}

var allowedImageSizes = []imageSize{
	{Width: 3840, Height: 1420},
	{Width: 900, Height: 1000},
	{Width: 1400, Height: 1866},
	{Width: 1866, Height: 1400},
	{Width: 1840, Height: 520},
	{Width: 2500, Height: 1666},
}

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

func LoadInterfaceAsJSON(source interface{}, target interface{}) error {
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

func IsCorrectImageSize(asset *contentful.AssetNoLocale) bool {
	image := GetAssetImage(asset)
	if image == nil {
		return false
	}
	for _, allowedImageSize := range allowedImageSizes {
		if image.Height == allowedImageSize.Height && image.Width == allowedImageSize.Width {
			return true
		}
	}
	return false
}

func IsLandscapeRatio(asset *contentful.AssetNoLocale) bool {
	aspectRatio, err := GetAspectRatio(asset)
	if err != nil {
		return false
	}
	return aspectRatio >= 1.00
}

func IsPortraitRatio(asset *contentful.AssetNoLocale) bool {
	aspectRatio, err := GetAspectRatio(asset)
	if err != nil {
		return false
	}
	return aspectRatio <= 1.00
}
