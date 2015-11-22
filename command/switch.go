package command

import (
	"github.com/S-Shimotori/ninjin/model"
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SwitchCommand struct {
	Meta
}

func (c *SwitchCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Printf(FailInGetVersion)
		return 1
	}

	xcodeList, xcodeError := function.GetXcodeList(function.ApplicationsPath)
	if xcodeError != nil {
		fmt.Printf(FailInMakingAListOfXcodes)
		return 1
	}

	if model.IsShortVersion(args[0]) {
		v, _ := model.NewShortVersion(args[0])

		for i := len(xcodeList) - 1; i >= 0; i-- {
			if model.EqualsForShortVersion(xcodeList[i].Version.Short, v) {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcodeList[i].AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcodeList[i].Version.Short), model.GetProductBuildVersionInString(xcodeList[i].Version.ProductBuild))
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

		for _, xcode := range xcodeList {
			if model.EqualsForProductBuildVersion(xcode.Version.ProductBuild, v) {
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
