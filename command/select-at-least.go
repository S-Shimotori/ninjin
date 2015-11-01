package command

import (
	"strings"
)

type SelectAtLeastCommand struct {
	Meta
}

func (c *SelectAtLeastCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SelectAtLeastCommand) Synopsis() string {
	return "select Xcode (at least [version])"
}

func (c *SelectAtLeastCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
