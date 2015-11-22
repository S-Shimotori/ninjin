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
		fmt.Printf(FailInGetVersion)
		return 1
	}
	xcodeLists, xcodeError := function.GetXcodeList(function.ApplicationsPath)
	if xcodeError != nil {
		fmt.Printf(FailInMakingAListOfXcodes)
		return 1
	}
	if model.IsShortVersion(args[0]) {
		v, _ := model.NewShortVersion(args[0])
		for _, xcode := range xcodeLists {
			// TODO: how to match Xcode's version
			if xcode.Version.Short == v {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcode.AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcode.Version.Short), model.GetProductBuildVersionInString(xcode.Version.ProductBuild))
					return 0
				} else {
					fmt.Printf(FailInSwitching, model.GetShortVersionInString(v))
					return 1
				}
			}
		}
		fmt.Printf(FailInFindingXcode, model.GetShortVersionInString(v))
	} else if model.IsProductBuildVersion(args[0]) {
		v, _ := model.NewProductBuildVersion(args[0])
		for _, xcode := range xcodeLists {
			// TODO: how to match Xcode's version
			if xcode.Version.ProductBuild == v {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcode.AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcode.Version.Short), model.GetProductBuildVersionInString(xcode.Version.ProductBuild))
					return 0
				} else {
					fmt.Printf(FailInSwitching, model.GetProductBuildVersionInString(v))
					return 1
				}
			}
		}
		fmt.Printf(FailInFindingXcode, model.GetProductBuildVersionInString(v))
	} else {
		fmt.Printf(FailInGetVersion)
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
