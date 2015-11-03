package model

import (
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
			xcode := Xcode{
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
