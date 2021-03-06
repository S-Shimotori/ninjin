package command

import (
	"github.com/S-Shimotori/ninjin/model"
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
	"os"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	xcodeList, xcodeError := function.GetXcodeList(function.ApplicationsPath)
	if xcodeError != nil {
		fmt.Printf(FailInMakingAListOfXcodes)
		os.Exit(1)
	}

	activeDeveloperDirectoryPath, _ := function.GetActiveDeveloperDirectoryPath()
	for _, xcode := range xcodeList {
		if activeDeveloperDirectoryPath == xcode.AppPath + function.PathToDeveloperDirectoryPath {
			fmt.Printf("* ")
		} else {
			fmt.Printf("  ")
		}

		fmt.Printf("%s (%s %s)\n", xcode.AppName, model.GetShortVersionInString(xcode.Version.Short), model.GetProductBuildVersionInString(xcode.Version.ProductBuild))
	}
	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List Xcodes in /Applications"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
