package model

import (
	"strings"
	"errors"
	"strconv"
)

type shortVersion struct {
	major int
	minor int
	patch int
}

func NewShortVersion(str string) (shortVersion, error) {
	var short shortVersion

	version := strings.Split(str, ".")

	if len(version) < 2 {
		return short, errors.New("invalid version")
	}

	major, majorError := strconv.Atoi(version[0])
	if majorError != nil {
		return short, errors.New("invalid major version")
	}

	minor, minorError := strconv.Atoi(version[1])
	if minorError != nil {
		return short, errors.New("invalid minor version")
	}

	if len(version) > 2 {
		patch, patchError := strconv.Atoi(version[2])
		if patchError == nil {
			return shortVersion{major, minor, patch}, nil
		}
	}

	return shortVersion{major, minor, 0}, nil
}

func LessForShortVersion(s0, s1 shortVersion) bool {
	if s0.major != s1.major {
		return s0.major < s1.major
	} else if s0.minor != s1.minor {
		return s0.minor < s1.minor
	} else {
		return s0.patch < s1.patch
	}
}

func EqualsForShortVersion(s0, s1 shortVersion) bool {
	return s0.major == s1.major && s0.minor == s1.minor && s0.patch == s1.patch
}

func IsShortVersion(str string) bool {
	_, shortVersionError := NewShortVersion(str)
	if shortVersionError == nil {
		return true
	} else {
		return false
	}
}

func GetExcessCompatibleShortVersion(str string) (shortVersion, error) {
	var short shortVersion
	if IsShortVersion(str) {
		versions := strings.Split(str, ".")
		if len(versions) > 2 {
			_, patchError := strconv.Atoi(versions[2])
			if patchError == nil {
				short, shortError := NewShortVersion(str)
				if shortError != nil {
					return short, shortError
				}
				short.minor += 1
				short.patch = 0
				return short, nil
			}
		}

		short, shortError := NewShortVersion(str)
		if shortError != nil {
			return short, shortError
		}

		short.major += 1
		short.minor = 0
		short.patch = 0
		return short, nil

	} else {
		return short, errors.New("not short version")
	}
}

func GetShortVersionInString(short shortVersion) string {
	main := strconv.Itoa(short.major) + "." + strconv.Itoa(short.minor)
	if short.patch == 0 {
		return main
	} else {
		return main + "." + strconv.Itoa(short.patch)
	}
}
