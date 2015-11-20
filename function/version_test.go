package function

import (
	"testing"
)

func TestGetCompatibleVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected string
	}{
		{"0", "1"},
		{"0.0", "1"},
		{"0.0.0", "0.1"},
		{"7A1001", "8"},
	}

	for _, c := range testCases {
		result := GetExcessCompatibleVersion(c.str)
		if result != c.expected {
			t.Errorf("got %v\nwant %v", result, c.expected)
		}
	}
}
