package model

import (
	"testing"
)

func TestIsVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected bool
	}{
		{"10", true},
		{"10.0", true},
		{"10.0.0", true},
		{"10.10", true},
		{"10", true},
		{".", false},
		{"10.", false},
		{"10..0", false},
		{"10.0.", false},
		{"00", false},
		{"10.00", false},
	}

	for _, c := range testCases {
		result := IsShortVersion(c.str)
		if result != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, result, c.expected)
		}
	}
}

func TestIsProductBuildVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected bool
	}{
		{"7A1001", true},
		{"7B91b", true},
		{"7C46l", true},
	}

	for _, c := range testCases {
		result := IsProductBuildVersion(c.str)
		if result != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, result, c.expected)
		}
	}
}
