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
		"select": func() (cli.Command, error) {
			return &command.SelectCommand{
				Meta: *meta,
			}, nil
		},
		"select-at-least": func() (cli.Command, error) {
			return &command.SelectAtLeastCommand{
				Meta: *meta,
			}, nil
		},
		"select-compatible": func() (cli.Command, error) {
			return &command.SelectCompatibleCommand{
				Meta: *meta,
			}, nil
		},
		"select-latest": func() (cli.Command, error) {
			return &command.SelectLatestCommand{
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
