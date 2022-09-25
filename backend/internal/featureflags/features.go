package featureflags

import "strings"

const (
	FeatureImgProxy         Feature = "image_proxy"
	FeatureImgSocialPreview Feature = "image_social_preview"
	FeatureReadability      Feature = "readability"
)

type Feature string

func (f Feature) String() string {
	return "FEATURE_FLAG_" + strings.ToUpper(string(f))
}
