package model

import (
	"testing"
	"reflect"
	"math/rand"
)

var versions = []version{
	newVersionWithoutError("7.0", "7A176x"),
	newVersionWithoutError("7.0", "7A192o"),
	newVersionWithoutError("7.0", "7A218"),
	newVersionWithoutError("7.0.1", "7A1001"),
	newVersionWithoutError("7.1", "7B75"),
	newVersionWithoutError("7.1", "7B85"),
	newVersionWithoutError("7.1", "7B91b"),
	newVersionWithoutError("7.1.1", "7B1005"),
	newVersionWithoutError("7.2", "7C46l"),
	newVersionWithoutError("7.2", "7C46t"),
}

func TestSort(t *testing.T) {
	ascending := make([]version, len(versions))
	copy(ascending, versions)

	descending := make([]version, len(versions))
	copy(descending, versions)
	for i, j := 0, len(descending) - 1; i < j; i, j = i + 1, j - 1 {
		descending[i], descending[j] = descending[j], descending[i]
	}
	var random []version = make([]version, len(versions))
	perm := rand.Perm(len(random))
	for i, v := range perm {
		random[v] = versions[i]
	}

	testList := [][]version{
		ascending,
		descending,
		random,
	}

	for _, array := range testList {
		actualXcodeList := XcodeSlice{}

		for _, version := range array {
			xcode := Xcode{
				AppPath: "",
				AppName: "",
				Version: version,
			}
			actualXcodeList = append(actualXcodeList, xcode)
		}

		actualXcodeList.Sort()
		actualVersionList := []version{}
		for _, xcode := range actualXcodeList {
			actualVersionList = append(actualVersionList, xcode.Version)
		}

		if !reflect.DeepEqual(actualVersionList[:], versions) {
			t.Errorf("got %v\nwant %v", actualVersionList[:], versions)
		}
	}
}
