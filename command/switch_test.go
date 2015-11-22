package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSwitchCommand_implement(t *testing.T) {
	var _ cli.Command = &SwitchCommand{}
}
