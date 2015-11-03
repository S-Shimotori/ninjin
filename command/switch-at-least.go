package command

import (
	"strings"
)

type SwitchAtLeastCommand struct {
	Meta
}

func (c *SwitchAtLeastCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SwitchAtLeastCommand) Synopsis() string {
	return "switch Xcode (at least [version])"
}

func (c *SwitchAtLeastCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
