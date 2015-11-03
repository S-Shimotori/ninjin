package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSwitchAtLeastCommand_implement(t *testing.T) {
	var _ cli.Command = &SwitchAtLeastCommand{}
}
