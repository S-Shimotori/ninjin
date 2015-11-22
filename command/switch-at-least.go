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
	// Write your code here
	if len(args) == 0 {
		fmt.Printf(FailInGetVersion)
		return 1
	}

	//TODO: ProductBuildVersion
	short, shortError := model.NewShortVersion(args[0])
	if shortError != nil {
		fmt.Printf(FailInGetVersion)
		return 1
	}
	xcodeLists, xcodeError := function.GetXcodeList(function.ApplicationsPath)
	if xcodeError != nil {
		fmt.Printf(FailInMakingAListOfXcodes)
		return 1
	}

	for i := len(xcodeLists) - 1; i >= 0; i-- {
		//TODO: イコールが入ってない?
		if model.LessForShortVersion(short, xcodeLists[i].Version.Short) {
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
	fmt.Printf(FailInFindingXcodeAtLeast, model.GetShortVersionInString(short))
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
