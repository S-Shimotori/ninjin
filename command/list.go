package command

import (
	"github.com/S-Shimotori/ninjin/function"
	"strings"
	"fmt"
	"os"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	// Write your code here
	xcodeLists, xcodesError := function.ListXcodes(function.ApplicationsPath)
	if xcodesError != nil {
		fmt.Println(xcodesError)
		os.Exit(1)
	}

	activeDeveloperDirectoryPath, _ := function.GetActiveDeveloperDirectoryPath()
	for _, xcode := range xcodeLists {
		if activeDeveloperDirectoryPath == xcode.AppPath + function.PathToDeveloperDirectoryPath {
			fmt.Printf("* ")
		} else {
			fmt.Printf("  ")
		}

		fmt.Printf("%s (%s %s)\n", xcode.AppName, xcode.ShortVersion, xcode.ProductBuildVersion)
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
