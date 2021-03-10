package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/spf13/cobra"
)

func New(config contract.ConfigAccessor) Module {
	return Module{config: config}
}

func Providers() di.Deps {
	return []interface{}{}
}

type Module struct {
	config contract.ConfigAccessor
}

func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewVersionCommand(m.config),
	)
}
