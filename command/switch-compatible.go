package command

import (
	"strings"
	"github.com/S-Shimotori/ninjin/model"
	"github.com/S-Shimotori/ninjin/function"
	"fmt"
)

type SwitchCompatibleCommand struct {
	Meta
}

func (c *SwitchCompatibleCommand) Run(args []string) int {
	// Write your code here
	if len(args) == 0 || !model.IsShortVersion(args[0]) && !model.IsProductBuildVersion(args[0]) {
		fmt.Printf(FailInGetVersion)
		return 1
	}

	if model.IsShortVersion(args[0]) {
		excess, excessError := model.GetExcessCompatibleShortVersion(args[0])
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
			//イコールの判定は？
			//7.2しかないときに7.3って渡されたら?
			if model.LessForShortVersion(xcodeLists[i].Version.Short, excess) {
				_, execError := function.ExecXcodeSelectSwitchOutput(xcodeLists[i].AppPath + function.PathToDeveloperDirectoryPath)
				if execError == nil {
					fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcodeLists[i].Version.Short), model.GetProductBuildVersionInString(xcodeLists[i].Version.ProductBuild))
					return 0
				} else {
					fmt.Printf(FailInSwitching, model.GetShortVersionInString(xcodeLists[i].Version.Short))
					return 1
				}
			}
		}

		fmt.Printf(FailInFindingXcodeCompatible, model.GetShortVersionInString(excess))
		return 1
	} else if model.IsProductBuildVersion(args[0]) {
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
			//イコールの判定は？
			if model.LessForProductBuildVersion(xcodeLists[i].Version.ProductBuild, excess) {
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

		fmt.Printf(FailInFindingXcodeCompatible, model.GetProductBuildVersionInString(excess))
		return 1
	} else {
		fmt.Printf(FailInGetVersion)
		return 1
	}

}

func (c *SwitchCompatibleCommand) Synopsis() string {
	return "switch Xcode (compatible with [version])"
}

func (c *SwitchCompatibleCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
