package command

import (
	"github.com/S-Shimotori/ninjin/model"
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SwitchAtLeastCommand struct {
	Meta
}

func (c *SwitchAtLeastCommand) Run(args []string) int {
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
			if model.EqualsForShortVersion(v, xcodeList[i].Version.Short) || model.LessForShortVersion(v, xcodeList[i].Version.Short) {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcodeList[i].AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcodeList[i].Version.Short), model.GetProductBuildVersionInString(xcodeList[i].Version.ProductBuild))
					return 0
				} else {
					fmt.Printf(FailInSwitching, model.GetShortVersionInString(xcodeList[i].Version.Short))
					return 1
				}
			}
		}
		fmt.Printf(FailInFindingXcodeAtLeast, model.GetShortVersionInString(v))

	} else if model.IsProductBuildVersion(args[0]) {
		v, _ := model.NewProductBuildVersion(args[0])

		for i := len(xcodeList) - 1; i >= 0; i-- {
			if model.EqualsForProductBuildVersion(v, xcodeList[i].Version.ProductBuild) || model.LessForProductBuildVersion(v, xcodeList[i].Version.ProductBuild) {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcodeList[i].AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcodeList[i].Version.Short), model.GetProductBuildVersionInString(xcodeList[i].Version.ProductBuild))
					return 0
				} else {
					fmt.Printf(FailInSwitching, model.GetProductBuildVersionInString(xcodeList[i].Version.ProductBuild))
					return 1
				}
			}
		}
		fmt.Printf(FailInFindingXcodeAtLeast, model.GetProductBuildVersionInString(v))

	} else {
		fmt.Printf(FailInGetVersion)
	}
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
