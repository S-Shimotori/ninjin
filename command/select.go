package command

import (
	"strings"
)

type SelectCommand struct {
	Meta
}

func (c *SelectCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SelectCommand) Synopsis() string {
	return "select Xcode (exactly [version])"
}

func (c *SelectCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
