package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSelectCommand_implement(t *testing.T) {
	var _ cli.Command = &SelectCommand{}
}
