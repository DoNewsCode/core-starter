package app

import (
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/nfangxu/core-skeleton/app/commands"
	"github.com/spf13/cobra"
)

func New(config contract.ConfigAccessor) Kernel {
	return Kernel{config: config}
}

func Providers() di.Deps {
	return []interface{}{}
}

type Kernel struct {
	config contract.ConfigAccessor
}

func (k Kernel) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewVersionCommand(k.config),
	)
}
