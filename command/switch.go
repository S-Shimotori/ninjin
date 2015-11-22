package command

import (
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
	"github.com/S-Shimotori/ninjin/model"
)

type SwitchCommand struct {
	Meta
}

func (c *SwitchCommand) Run(args []string) int {
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
	if model.IsShortVersion(args[0]) {
		v, _ := model.NewShortVersion(args[0])
		for _, xcode := range xcodeLists {
			// TODO: how to match Xcode's version
			if xcode.Version.Short == v {
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
		fmt.Printf("can't find Xcode(version %s)\n", model.GetShortVersionInString(v))
	} else if model.IsProductBuildVersion(args[0]) {
		v, _ := model.NewProductBuildVersion(args[0])
		for _, xcode := range xcodeLists {
			// TODO: how to match Xcode's version
			if xcode.Version.ProductBuild == v {
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
		fmt.Printf("can't find Xcode(version %s)\n", model.GetProductBuildVersionInString(v))
	} else {
		fmt.Println("This command requires Xcode's valid version.")
	}
	return 1
}

func (c *SwitchCommand) Synopsis() string {
	return "switch Xcode (exactly [version])"
}

func (c *SwitchCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
