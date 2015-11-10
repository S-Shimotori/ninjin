package function

import (
	"regexp"
	"strings"
	"strconv"
)

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

func GetExtraCompatibleVersion(str string) string {
	if IsShortVersion(str) {
		nums := strings.Split(str, ".")

		if len(nums) == 1 {
			nums = append(nums, "0")
		}

		last2, atoiError := strconv.Atoi(nums[len(nums) - 2])
		if atoiError != nil {
			return ""
		}
		nums[len(nums) - 2] = strconv.Itoa(last2 + 1)
		return strings.Join(nums[0:(len(nums) - 1)], ".")

	} else if IsProductBuildVersion(str) {
		return strconv.Itoa(int([]rune(str)[0] - '0') + 1)

	} else {
		return ""
	}
}
