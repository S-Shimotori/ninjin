package model

import (
	"regexp"
	"strings"
	"strconv"
)

func less(ver1, ver2 string) bool {
	if IsVersion(ver1) && IsVersion(ver2) == false {
		return true
	}

	nums1 := strings.Split(ver1, ".")
	nums2 := strings.Split(ver2, ".")

	if len(nums1) > len(nums2) {
		diff := len(nums1) - len(nums2)
		for i := 0; i < diff; i++ {
			nums2 = append(nums2, "0")
		}
	} else if len(nums1) < len(nums2) {
		diff := len(nums2) - len(nums1)
		for i := 0; i < diff; i++ {
			nums1 = append(nums1, "0")
		}
	}

	for i, _ := range nums1 {
		n1, error1 := strconv.Atoi(nums1[i])
		n2, error2 := strconv.Atoi(nums2[i])
		if error1 != nil || error2 != nil {
			return true
		}

		if n1 == n2 {
			continue
		} else {
			return n1 < n2
		}
	}
	return true
}

func IsVersion(str string) bool {
	pattern := `^(0|[1-9]+[0-9]*)(\.(0|[1-9]+[0-9]*))*$`
	result, error := regexp.MatchString(pattern, str)

	if error != nil {
		return false
	}

	return result
}
