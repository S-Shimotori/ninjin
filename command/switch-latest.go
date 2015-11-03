package command

import (
	"strings"
)

type SwitchLatestCommand struct {
	Meta
}

func (c *SwitchLatestCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SwitchLatestCommand) Synopsis() string {
	return "switch Xcode (latest version)"
}

func (c *SwitchLatestCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
