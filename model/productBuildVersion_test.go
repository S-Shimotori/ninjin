package model

import (
	"testing"
)

func TestNewProductBuildVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected productBuildVersion
	}{
		{"7", productBuildVersion{}},
		{"7A", productBuildVersion{7, "A", ""}},
		{"7A0", productBuildVersion{7, "A", "0"}},

		{"10", productBuildVersion{}},
		{"10A", productBuildVersion{10, "A", ""}},
		{"10A0", productBuildVersion{10, "A", "0"}},

		{"aA0", productBuildVersion{}},
		{"7a0", productBuildVersion{}},
		{"7A-", productBuildVersion{7, "A", ""}},

		{"A0", productBuildVersion{}},
	}

	for _, c := range testCases {
		productBuild, productBuildError := NewProductBuildVersion(c.str)
		if productBuildError != nil && (c.expected != productBuildVersion{}) {
			t.Errorf("got %v(productBuildVersion)\nwant %v", c.str, c.expected)
		} else if productBuild != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, productBuild, c.expected)
		}
	}
}

func TestIsProductBuildVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected bool
	}{
		{"7", false},
		{"7A", true},
		{"7A0", true},

		{"10", false},
		{"10A", true},
		{"10A0", true},

		{"aA0", false},
		{"7a0", false},
		{"7A-", true},

		{"A0", false},
	}

	for _, c := range testCases {
		result := IsProductBuildVersion(c.str)
		if result != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, result, c.expected)
		}
	}
}

func TestGetExcessCompatibleProductBuildVersion(t *testing.T) {
	var testCases = []struct {
		str string
		expected productBuildVersion
	}{
		{"7", productBuildVersion{}},
		{"7A", productBuildVersion{8, "A", ""}},
		{"7A0", productBuildVersion{7, "B", ""}},

		{"10", productBuildVersion{}},
		{"10A", productBuildVersion{11, "A", ""}},
		{"10A0", productBuildVersion{10, "B", ""}},

		{"aA0", productBuildVersion{}},
		{"7a0", productBuildVersion{}},
		{"7A-", productBuildVersion{8, "A", ""}},

		{"A0", productBuildVersion{}},
	}

	for _, c := range testCases {
		short, shortError := GetExcessCompatibleProductBuildVersion(c.str)
		if shortError != nil && (c.expected != productBuildVersion{}) {
			t.Errorf("got %v(productBuildVersion)\nwant %v", c.str, c.expected)
		} else if short != c.expected {
			t.Errorf("got %v(%v)\nwant %v", c.str, short, c.expected)
		}
	}
}
