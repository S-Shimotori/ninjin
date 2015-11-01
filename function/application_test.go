package function

import (
	"testing"
)

const plistsDirectory string = "../plists/"
func TestGetVersions(t *testing.T) {
	var testCases = []struct {
		appName, expectedShortVersion, expectedBuildVersion string
	}{
		{"Xcode.app", "7.2", "7C46l"},
	}

	for _, c := range testCases {
		actualShortVersion, actualBuildVersion := GetVersions(plistsDirectory + c.appName)
		if actualShortVersion != c.expectedShortVersion {
			t.Errorf("got %v\nwant %v", actualShortVersion, c.expectedShortVersion)
		}
		if actualBuildVersion != c.expectedBuildVersion {
			t.Errorf("got %v\nwant %v", actualBuildVersion, c.expectedBuildVersion)
		}
	}
}
