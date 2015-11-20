package command

import (
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SwitchAtLeastCommand struct {
	Meta
}

func (c *SwitchAtLeastCommand) Run(args []string) int {
	// Write your code here
	if len(args) == 0 || !function.IsShortVersion(args[0]) {
		fmt.Println("This command requires Xcode's version.")
		return 1
	}
	xcodeLists, xcodesError := function.ListXcodes(function.ApplicationsPath)
	if xcodesError != nil {
		fmt.Println(xcodesError)
		return 1
	}

	for i := len(xcodeLists) - 1; i >= 0; i-- {
		if function.Less(args[0], xcodeLists[i].Version.Short) {
			_, execError := function.ExecXcodeSelectSwitchOutput(xcodeLists[i].AppPath + function.PathToDeveloperDirectoryPath)
			if execError == nil {
				fmt.Printf("success.\n")
				return 0
			} else {
				fmt.Println(execError)
				return 1
			}
		}
	}
	fmt.Printf("can't find Xcode(version %s)\n", args[0])
	return 1
}

func (c *SwitchAtLeastCommand) Synopsis() string {
	return "switch Xcode (at least [version])"
}

func (c *SwitchAtLeastCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
