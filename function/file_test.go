package function

import (
	"testing"
	"reflect"
)

func TestListFiles(t *testing.T) {
	var testCases = []struct {
		directoryPath  string
		expectedList []string
	}{
		{"../plists/Xcode.app/Contents/", []string{"Info.plist", "version.plist"}},
	}

	for _, c := range testCases {
		actualList, _ := listFiles(c.directoryPath)

		if !reflect.DeepEqual(actualList[:], c.expectedList[:]) {
			t.Errorf("got %v\nwant %v", actualList, c.expectedList)
		}
	}
}
