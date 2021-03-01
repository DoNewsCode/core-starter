package config

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core/srvhttp"
	"github.com/nfangxu/core-skeleton/app/commands"
	"github.com/nfangxu/core-skeleton/app/kernel"
	"github.com/spf13/cobra"
)

var (
	// Register providers
	Providers = []interface{}{
		//
	}

	// Register modules
	Modules = []interface{}{
		kernel.InjectKernel,
		core.NewServeModule,
		srvhttp.HealthCheckModule{},
	}

	// Register commands
	Commands = []func(c *core.C) *cobra.Command{
		commands.NewVersionCmd,
	}
)
