package command

import (
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SwitchLatestCommand struct {
	Meta
}

func (c *SwitchLatestCommand) Run(args []string) int {
	// Write your code here
	xcodeLists, xcodeError := function.GetXcodeList(function.ApplicationsPath)
	if xcodeError != nil {
		fmt.Println(xcodeError)
		return 1
	}
	if len(xcodeLists) == 0 {
		fmt.Printf("can't find Xcode\n")
		return 1
	}

	xcode := xcodeLists[len(xcodeLists) - 1]
	_, execError := function.ExecXcodeSelectSwitchOutput(xcode.AppPath + function.PathToDeveloperDirectoryPath)
	if execError == nil {
		fmt.Printf("success.\n")
		return 0
	} else {
		fmt.Println(execError)
		return 1
	}
}

func (c *SwitchLatestCommand) Synopsis() string {
	return "switch Xcode (latest version)"
}

func (c *SwitchLatestCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
