package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/spf13/cobra"
)

// New is the constructor of the app module.
func New(config contract.ConfigAccessor) Module {
	return Module{config: config}
}

// Providers registers the dependencies provided by the app module (so that other modules can consume them).
func Providers() di.Deps {
	return di.Deps{}
}

// Module is the struct representing the app Module.
type Module struct {
	config contract.ConfigAccessor
}

// ProvideCommand implements container.CommandProvider
func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewExampleCommand(m.config),
	)
}
