package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSelectAtLeastCommand_implement(t *testing.T) {
	var _ cli.Command = &SelectAtLeastCommand{}
}
