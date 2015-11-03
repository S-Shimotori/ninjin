package model

import (
	"sort"
)

type Xcode struct {
	AppPath string
	AppName string
	ShortVersion string
	ProductBuildVersion string
}

type XcodeSlice []Xcode
func (p XcodeSlice) Len() int {
	return len(p)
}
func (p XcodeSlice) Less(i, j int) bool {
	// TODO: compare version
	if p[i].ShortVersion != p[j].ShortVersion {
		return p[i].ShortVersion < p[j].ShortVersion
	} else if p[i].ProductBuildVersion != p[j].ProductBuildVersion {
		return p[i].ProductBuildVersion < p[j].ProductBuildVersion
	} else {
		return true
	}
}
func (p XcodeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p XcodeSlice) Sort() {
	sort.Sort(p)
}
