package command

import (
	"github.com/S-Shimotori/ninjin/model"
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
)

type SwitchLatestCommand struct {
	Meta
}

func (c *SwitchLatestCommand) Run(args []string) int {
	xcodeList, xcodeError := function.GetXcodeList(function.ApplicationsPath)
	if xcodeError != nil {
		fmt.Printf(FailInMakingAListOfXcodes)
		return 1
	}

	if len(xcodeList) == 0 {
		fmt.Printf(FailInFindingXcodeLatest)
		return 1
	}

	xcode := xcodeList[len(xcodeList) - 1]
	_, execError := function.ExecXcodeSelectSwitchOutput(xcode.AppPath + function.PathToDeveloperDirectoryPath)
	if execError == nil {
		fmt.Printf(SucceedInSwitching, model.GetShortVersionInString(xcode.Version.Short), model.GetProductBuildVersionInString(xcode.Version.ProductBuild))
		return 0
	} else {
		fmt.Printf(FailInSwitching, model.GetShortVersionInString(xcode.Version.Short) + " " + model.GetProductBuildVersionInString(xcode.Version.ProductBuild))
		return 1
	}

}

func (c *SwitchLatestCommand) Synopsis() string {
	return "switch Xcode (latest version)"
}

func (c *SwitchLatestCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
