package model

import (
	"testing"
)

func TestLess(t *testing.T) {
	var testCases = []struct {
		ver1 string
		ver2 string
		expected bool
	}{
		{"10", "10.0", true},
		{"10", "10.1", true},
		{"10", "10.0.1", true},
		{"10", "11", true},
		{"10", "9", false},
		{"10", "9.0", false},
		{"10", "9.1", false},
		{"10.0", "9", false},
		{"10.1", "10.2", true},
	}

	for _, c := range testCases {
		result := less(c.ver1, c.ver2)
		if result != c.expected {
			t.Errorf("got %v %v(%v)\nwant %v", c.ver1, c.ver2, result, c.expected)
		}
	}
}

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
		result := IsVersion(c.str)
		if result != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, result, c.expected)
		}
	}
}
