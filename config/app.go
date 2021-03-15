package config

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core-starter/app"
	"github.com/DoNewsCode/core/config"
	"github.com/DoNewsCode/core/di"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/DoNewsCode/core/srvhttp"
)

var (
	// Register providers
	Providers = []di.Deps{
		app.Providers(),
		otgorm.Providers(),
	}

	// Register modules
	Modules = []interface{}{
		app.New,
		config.New,                  // config module
		core.NewServeModule,         // server module
		srvhttp.HealthCheckModule{}, // health check module (http demo)
		srvhttp.DocsModule{},        // docs
		otgorm.New,
	}
)
