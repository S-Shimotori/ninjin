package model

import (
	"testing"
)

func TestNewShortVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected shortVersion
	}{
		{"10", shortVersion{}},
		{"10.1", shortVersion{10, 1, 0}},
		{"10.1.1", shortVersion{10, 1, 1}},
		{"10.1.1.1", shortVersion{10, 1, 1}},

		{"a.1.1", shortVersion{}},
		{"10.a.1", shortVersion{}},
		{"10.1.a", shortVersion{10, 1, 0}},
		{"10.1.1.a", shortVersion{10, 1, 1}},

		{".", shortVersion{}},
		{".1.1", shortVersion{}},
		{"10..1", shortVersion{}},
		{"10.1.", shortVersion{10, 1, 0}},
	}

	for _, c := range testCases {
		short, shortError := NewShortVersion(c.str)
		if shortError != nil && (c.expected != shortVersion{}) {
			t.Errorf("got %v(shortVersion)\nwant %v", c.str, c.expected)
		} else if short != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, short, c.expected)
		}
	}
}

func TestGetExcessCompatibleShortVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected shortVersion
	}{
		{"10", shortVersion{}},
		{"10.1", shortVersion{11, 0, 0}},
		{"10.1.1", shortVersion{10, 2, 0}},
		{"10.1.1.1", shortVersion{10, 2, 0}},

		{"a.1.1", shortVersion{}},
		{"10.a.1", shortVersion{}},
		{"10.1.a", shortVersion{11, 0, 0}},
		{"10.1.1.a", shortVersion{10, 2, 0}},

		{".", shortVersion{}},
		{".1.1", shortVersion{}},
		{"10..1", shortVersion{}},
		{"10.1.", shortVersion{11, 0, 0}},
	}

	for _, c := range testCases {
		short, shortError := GetExcessCompatibleShortVersion(c.str)
		if shortError != nil && (c.expected != shortVersion{}) {
			t.Errorf("got %v(shortVersion)\nwant %v", c.str, c.expected)
		} else if short != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, short, c.expected)
		}
	}
}
