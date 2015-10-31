package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSelectCompatibleCommand_implement(t *testing.T) {
	var _ cli.Command = &SelectCompatibleCommand{}
}
