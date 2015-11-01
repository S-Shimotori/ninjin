package command

import (
	"strings"
)

type SelectLatestCommand struct {
	Meta
}

func (c *SelectLatestCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SelectLatestCommand) Synopsis() string {
	return "select Xcode (latest version)"
}

func (c *SelectLatestCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
