package main

import (
	"github.com/S-Shimotori/ninjin/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"switch": func() (cli.Command, error) {
			return &command.SwitchCommand{
				Meta: *meta,
			}, nil
		},
		"switch-at-least": func() (cli.Command, error) {
			return &command.SwitchAtLeastCommand{
				Meta: *meta,
			}, nil
		},
		"switch-compatible": func() (cli.Command, error) {
			return &command.SwitchCompatibleCommand{
				Meta: *meta,
			}, nil
		},
		"switch-latest": func() (cli.Command, error) {
			return &command.SwitchLatestCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
