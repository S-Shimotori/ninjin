package command

import (
	"strings"
)

type SelectCompatibleCommand struct {
	Meta
}

func (c *SelectCompatibleCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SelectCompatibleCommand) Synopsis() string {
	return "select Xcode (compatible with [version])"
}

func (c *SelectCompatibleCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
