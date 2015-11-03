package command

import (
	"strings"
)

type SwitchCompatibleCommand struct {
	Meta
}

func (c *SwitchCompatibleCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SwitchCompatibleCommand) Synopsis() string {
	return "switch Xcode (compatible with [version])"
}

func (c *SwitchCompatibleCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
