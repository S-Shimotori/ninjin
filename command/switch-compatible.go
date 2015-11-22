package command

import (
	"github.com/S-Shimotori/ninjin/model"
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SwitchCompatibleCommand struct {
	Meta
}

func (c *SwitchCompatibleCommand) Run(args []string) int {
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

		excess, excessError := model.GetExcessCompatibleShortVersion(args[0])
		if excessError != nil {
			fmt.Printf(FailInGetVersion)
			return 1
		}

		for i := len(xcodeList) - 1; i >= 0; i-- {
			if !model.EqualsForShortVersion(xcodeList[i].Version.Short, v) && model.LessForShortVersion(xcodeList[i].Version.Short, v) {
				break
			}
			if !model.EqualsForShortVersion(xcodeList[i].Version.Short, excess) && model.LessForShortVersion(xcodeList[i].Version.Short, excess) {
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
		fmt.Printf(FailInFindingXcodeCompatible, model.GetShortVersionInString(v))

	} else if model.IsProductBuildVersion(args[0]) {
		v, _ := model.NewProductBuildVersion(args[0])

		excess, excessError := model.GetExcessCompatibleProductBuildVersion(args[0])
		if excessError != nil {
			fmt.Printf(FailInGetVersion)
			return 1
		}

		xcodeLists, xcodeError := function.GetXcodeList(function.ApplicationsPath)
		if xcodeError != nil {
			fmt.Printf(FailInMakingAListOfXcodes)
			return 1
		}

		for i := len(xcodeLists) - 1; i >= 0; i-- {
			if !model.EqualsForProductBuildVersion(xcodeLists[i].Version.ProductBuild, v) && model.LessForProductBuildVersion(xcodeLists[i].Version.ProductBuild, v) {
				break
			}
			if !model.EqualsForProductBuildVersion(xcodeLists[i].Version.ProductBuild, excess) && model.LessForProductBuildVersion(xcodeLists[i].Version.ProductBuild, excess) {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcodeLists[i].AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcodeLists[i].Version.Short), model.GetProductBuildVersionInString(xcodeLists[i].Version.ProductBuild))
					return 0
				} else {
					fmt.Println(execError)
					return 1
				}
			}
		}
		fmt.Printf(FailInFindingXcodeCompatible, model.GetProductBuildVersionInString(v))

	} else {
		fmt.Printf(FailInGetVersion)
	}
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
