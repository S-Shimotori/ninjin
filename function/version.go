package function

import (
	"regexp"
)

func IsVersion(str string) bool {
	pattern := `^(0|[1-9]+[0-9]*)(\.(0|[1-9]+[0-9]*))*$`
	result, error := regexp.MatchString(pattern, str)

	if error != nil {
		return false
	}

	return result
}
