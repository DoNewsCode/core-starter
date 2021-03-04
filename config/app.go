package config

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core/config"
	"github.com/DoNewsCode/core/di"
	"github.com/DoNewsCode/core/srvhttp"
	"github.com/nfangxu/core-skeleton/app"
)

var (
	// Register providers
	Providers = []di.Deps{
		app.Providers(),
	}

	// Register modules
	Modules = []interface{}{
		app.New,
		config.New,                  // config module
		core.NewServeModule,         // server module
		srvhttp.HealthCheckModule{}, // health check module (http demo)
	}
)
