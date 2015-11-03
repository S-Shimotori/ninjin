package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSwitchCompatibleCommand_implement(t *testing.T) {
	var _ cli.Command = &SwitchCompatibleCommand{}
}
