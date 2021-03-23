package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/spf13/cobra"
)

// New Create a new module
func New(config contract.ConfigAccessor) Module {
	return Module{config: config}
}

// Providers Register the dependencies of this module
func Providers() di.Deps {
	return di.Deps{}
}

// Module is a main file
type Module struct {
	config contract.ConfigAccessor
}

// ProvideCommand Marks that this module provides some commands
func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewExampleCommand(m.config),
	)
}
