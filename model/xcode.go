package model

import (
	"sort"
)

type Xcode struct {
	AppPath string
	AppName string
	Version version
}

type XcodeSlice []Xcode

func (p XcodeSlice) Len() int {
	return len(p)
}

func (p XcodeSlice) Less(i, j int) bool {
	return Less(p[i].Version, p[j].Version)
}

func (p XcodeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p XcodeSlice) Sort() {
	sort.Sort(p)
}
