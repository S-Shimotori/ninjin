package command

import (
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SelectCommand struct {
	Meta
}

func (c *SelectCommand) Run(args []string) int {
	// Write your code here
	if len(args) == 0 {
		fmt.Println("This command requires Xcode's version.")
		return 1
	}
	xcodeLists, xcodesError := function.ListXcodes(function.ApplicationsPath)
	if xcodesError != nil {
		fmt.Println(xcodesError)
		return 1
	}

	for _, xcode := range xcodeLists {
		if xcode.ProductBuildVersion == args[0] || xcode.ShortVersion == args[0] {
			_, execError := function.ExecXcodeSelectSwitchOutput(xcode.AppPath + function.PathToDeveloperDirectoryPath)
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

func (c *SelectCommand) Synopsis() string {
	return "select Xcode (exactly [version])"
}

func (c *SelectCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
