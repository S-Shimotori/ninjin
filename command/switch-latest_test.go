package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSwitchLatestCommand_implement(t *testing.T) {
	var _ cli.Command = &SwitchLatestCommand{}
}
