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
		fmt.Errorf("%s", xcodesError)
		os.Exit(1)
	}

	for _, xcode := range xcodeLists {
		if function.IsActiveXcode(function.GenerateFullPathForFileInApplications(xcode)) {
			fmt.Printf("* ")
		} else {
			fmt.Printf("  ")
		}
		shortVersion, buildVersion := function.GetVersions(function.GenerateFullPathForFileInApplications(xcode))
		version := ""
		switch {
		case shortVersion != "" && buildVersion != "":
			version = shortVersion + " " + buildVersion
		case shortVersion != "" && buildVersion == "":
			version = shortVersion + " "
		case shortVersion == "" && buildVersion != "":
			version = " " + buildVersion
		}

		fmt.Printf("%s (%s)\n", xcode, version)
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
