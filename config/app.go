package config

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core/config"
	"github.com/DoNewsCode/core/di"
	"github.com/DoNewsCode/core/srvhttp"
	"github.com/nfangxu/core-skeleton/app/commands"
	"github.com/nfangxu/core-skeleton/app/kernel"
	"github.com/spf13/cobra"
)

var (
	// Register providers
	Providers = []di.Deps{
		// otgorm.Providers(), // register gorm providers
	}

	// Register modules
	Modules = []interface{}{
		kernel.InjectKernel, // Application kernel
		config.New, // config module
		core.NewServeModule, // server module
		srvhttp.HealthCheckModule{}, // health check module (http demo)
		// otgorm.New, // gorm module
	}

	// Register commands
	Commands = []func(c *core.C) *cobra.Command{
		commands.NewVersionCmd, // version command
	}
)
