package command

import (
	"strings"
	"github.com/S-Shimotori/ninjin/function"
	"fmt"
)

type SwitchCompatibleCommand struct {
	Meta
}

func (c *SwitchCompatibleCommand) Run(args []string) int {
	// Write your code here
	if len(args) == 0 || !function.IsShortVersion(args[0]) && !function.IsProductBuildVersion(args[0]) {
		fmt.Println("This command requires Xcode's version.")
		return 1
	}
	extraCompatibleVersion := function.GetExcessCompatibleVersion(args[0])

	xcodeLists, xcodesError := function.ListXcodes(function.ApplicationsPath)
	if xcodesError != nil {
		fmt.Println(xcodesError)
		return 1
	}
	fmt.Printf("%s\n", extraCompatibleVersion)
	for i := len(xcodeLists) - 1; i >= 0; i-- {
		fmt.Println(xcodeLists[i].Version.Short)
		//TODO: イコールの排除
		if (function.IsShortVersion(args[0]) && function.Less(args[0], xcodeLists[i].Version.Short) ||
			function.IsProductBuildVersion(args[0]) && function.Less(args[0], xcodeLists[i].Version.ProductBuild)) &&
		function.Less(xcodeLists[i].Version.Short, extraCompatibleVersion) {
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

func (c *SwitchCompatibleCommand) Synopsis() string {
	return "switch Xcode (compatible with [version])"
}

func (c *SwitchCompatibleCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
