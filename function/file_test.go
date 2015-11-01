package function

import (
	"testing"
	"reflect"
)

func TestListApplications(t *testing.T) {
	var testCases = []struct {
		directoryPath  string
		expectedList []string
	}{
		{"../plists/", []string{"Xcode.app"}},
	}

	for _, c := range testCases {
		actualList, _ := ListApplications(c.directoryPath)

		if !reflect.DeepEqual(actualList[:], c.expectedList[:]) {
			t.Errorf("got %v\nwant %v", actualList, c.expectedList)
		}
	}
}
