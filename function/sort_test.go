package function

import (
	"github.com/S-Shimotori/ninjin/model"
	"testing"
	"reflect"
)

func TestSort(t *testing.T) {
	var testCases = []struct {
		testList [][]string
		expectedList [][]string
	}{
		{
			[][]string{{"7.2", "7C461"}, {"7.0.1", "7A1001"}, {"7.1", "7B91b"}},
			[][]string{{"7.0.1", "7A1001"}, {"7.1", "7B91b"}, {"7.2", "7C461"}},
		},
		{
			[][]string{{"", "7C461"}, {"", "7A1001"}, {"", "7B91b"}},
			[][]string{{"", "7A1001"}, {"", "7B91b"}, {"", "7C461"}},
		},
		{
			[][]string{{"7.0.1", ""}, {"7", ""}, {"7.1", ""}},
			[][]string{{"7", ""}, {"7.0.1", ""}, {"7.1", ""}},
		},
	}

	for _, arrays := range testCases {
		actualXcodeList := XcodeSlice{}
		actualStringList := [][]string{}

		for _, versions := range arrays.testList {
			xcode := model.Xcode{
				AppPath: "",
				AppName: "",
				ShortVersion: versions[0],
				ProductBuildVersion: versions[1],
			}
			actualXcodeList = append(actualXcodeList, xcode)
		}
		actualXcodeList.Sort()

		for _, xcode := range actualXcodeList {
			actualStringList = append(actualStringList, []string{xcode.ShortVersion, xcode.ProductBuildVersion})
		}
		if !reflect.DeepEqual(actualStringList[:], arrays.expectedList[:]) {
			t.Errorf("got %v\nwant %v", actualStringList, arrays.expectedList)
		}
	}
}

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
		result := Less(c.ver1, c.ver2)
		if result != c.expected {
			t.Errorf("got %v %v(%v)\nwant %v", c.ver1, c.ver2, result, c.expected)
		}
	}
}
