package function

import (
	"github.com/S-Shimotori/ninjin/model"
	"strings"
	"strconv"
)

func GetExcessCompatibleVersion(str string) string {
	if model.IsShortVersion(str) {
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

	} else if model.IsProductBuildVersion(str) {
		return strconv.Itoa(int([]rune(str)[0] - '0') + 1)

	} else {
		return ""
	}
}
