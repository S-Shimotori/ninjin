package model

import (
	"regexp"
	"errors"
)

type Version struct {
	Short string
	ProductBuild string
}

func IsShortVersion(str string) bool {
	pattern := `^(0|[1-9]+[0-9]*)(\.(0|[1-9]+[0-9]*))*$`
	result, matchError := regexp.MatchString(pattern, str)

	if matchError != nil {
		return false
	}

	return result
}

func IsProductBuildVersion(str string) bool {
	pattern := `^[1-9][0-9]*[A-Z][0-9a-z]+$`
	result, matchError := regexp.MatchString(pattern, str)

	if matchError != nil {
		return false
	}

	return result
}
func NewVersion(shortVersion string, productBuildVersion string) (Version, error) {
	if IsShortVersion(shortVersion) && IsProductBuildVersion(productBuildVersion) {
		return Version{Short: shortVersion, ProductBuild: productBuildVersion}, nil
	} else {
		return Version{}, errors.New("invalid version")
	}
}

type Xcode struct {
	AppPath string
	AppName string
	Version Version
}
