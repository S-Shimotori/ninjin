package model

import (
	"regexp"
	"errors"
	"strconv"
)

type productBuildVersion struct {
	major int
	minor string
	patch string
}

func NewProductBuildVersion(str string) (productBuildVersion, error) {
	var productBuild productBuildVersion

	if !IsProductBuildVersion(str) {
		return productBuild, errors.New("invalid version")
	}

	mainRegexp := regexp.MustCompile("^([1-9][0-9]*)([A-Z])")
	main := mainRegexp.FindSubmatch([]byte(str))
	if len(main) < 3 {
		return productBuild, errors.New("invalid version")
	}

	major, majorError := strconv.Atoi(string(main[1]))
	if majorError != nil {
		return productBuild, errors.New("invalid major version")
	}

	minor := string(main[2])

	patchRegexp, _ := regexp.Compile("^[1-9][0-9]*[A-Z]([0-9a-z]+)$")
	patch := patchRegexp.FindSubmatch([]byte(str))
	if len(patch) < 2 {
		return productBuildVersion{major, minor, ""}, nil
	} else {
		return productBuildVersion{major, minor, string(patch[1])}, nil
	}
}

func LessForProductBuildVersion(pb0, pb1 productBuildVersion) bool {
	if pb0.major != pb1.major {
		return pb0.major < pb1.major
	} else if pb0.minor != pb1.minor {
		return pb0.minor < pb1.minor
	} else {
		return pb0.patch < pb1.patch
	}
}

func EqualsForProductBuildVersion(pb0, pb1 productBuildVersion) bool {
	return pb0.major == pb1.major && pb0.minor == pb1.minor && pb0.patch == pb1.patch
}

func IsProductBuildVersion(str string) bool {
	pattern := `^[1-9][0-9]*[A-Z]([0-9a-z]*)$`
	result, matchError := regexp.MatchString(pattern, str)

	if matchError != nil {
		return false
	}

	return result
}

func GetExcessCompatibleProductBuildVersion(str string) (productBuildVersion, error) {
	var productBuild productBuildVersion
	if IsProductBuildVersion(str) {
		mainRegexp := regexp.MustCompile("^([1-9][0-9]*)([A-Z])")
		main := mainRegexp.FindSubmatch([]byte(str))
		if len(main) < 3 {
			return productBuild, errors.New("invalid version")
		}

		major, majorError := strconv.Atoi(string(main[1]))
		if majorError != nil {
			return productBuild, errors.New("invalid major version")
		}

		minor := string(main[2])

		patchRegexp, _ := regexp.Compile("^[1-9][0-9]*[A-Z]([0-9a-z]+)$")
		patch := patchRegexp.FindSubmatch([]byte(str))
		if len(patch) < 2 {
			return productBuildVersion{major + 1, "A", ""}, nil
		}

		return productBuildVersion{major, string([]byte(minor)[0] + 1), ""}, nil
	} else {
		return productBuild, errors.New("not product build version")
	}
}

func GetProductBuildVersionInString(short productBuildVersion) string {
	return strconv.Itoa(short.major) + short.minor + short.patch
}
