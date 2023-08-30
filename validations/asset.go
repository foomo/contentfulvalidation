package validations

import (
	"github.com/foomo/contentful"
	"github.com/foomo/contentfulvalidation/utils"
)

func IsLandscapeRatio(asset *contentful.AssetNoLocale) bool {
	aspectRatio, err := utils.GetAspectRatio(asset)
	if err != nil {
		return false
	}
	return aspectRatio >= 1.00
}

func IsPortraitRatio(asset *contentful.AssetNoLocale) bool {
	aspectRatio, err := utils.GetAspectRatio(asset)
	if err != nil {
		return false
	}
	return aspectRatio <= 1.00
}
