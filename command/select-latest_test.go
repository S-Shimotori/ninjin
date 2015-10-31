package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSelectLatestCommand_implement(t *testing.T) {
	var _ cli.Command = &SelectLatestCommand{}
}
