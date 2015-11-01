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
		shortVersion, buildVersion := function.GetVersions(function.GenerateFullPathForFileInApplications(xcode))
		version := ""
		switch {
		case shortVersion != "" && buildVersion != "":
			version = shortVersion + "\t" + buildVersion
		case shortVersion != "" && buildVersion == "":
			version = shortVersion + "\t"
		case shortVersion == "" && buildVersion != "":
			version = "\t" + buildVersion
		}

		fmt.Printf("%s\t(%s)\n", xcode, version)
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
