package function

import (
	"testing"
	"reflect"
)

const plistsDirectory string = "../plists/"
func TestGetVersions(t *testing.T) {
	var testCases = []struct {
		appName, expectedShortVersion, expectedBuildVersion string
	}{
		{"Xcode.app", "7.2", "7C46l"},
	}

	for _, c := range testCases {
		actualShortVersion, actualBuildVersion := getVersions(plistsDirectory + c.appName)
		if actualShortVersion != c.expectedShortVersion {
			t.Errorf("got %v\nwant %v", actualShortVersion, c.expectedShortVersion)
		}
		if actualBuildVersion != c.expectedBuildVersion {
			t.Errorf("got %v\nwant %v", actualBuildVersion, c.expectedBuildVersion)
		}
	}
}

func TestListApplications(t *testing.T) {
	var testCases = []struct {
		directoryPath  string
		expectedList []string
	}{
		{"../plists/", []string{"Xcode.app"}},
	}

	for _, c := range testCases {
		actualXcodeList, _ := GetXcodeList(c.directoryPath)
		actualStringList := []string{}
		for _, xcode := range actualXcodeList {
			actualStringList = append(actualStringList, xcode.AppName)
		}
		if !reflect.DeepEqual(actualStringList[:], c.expectedList[:]) {
			t.Errorf("got %v\nwant %v", actualStringList, c.expectedList)
		}
	}
}
