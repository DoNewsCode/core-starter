package config

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core-starter/app"
	"github.com/DoNewsCode/core-starter/docs"
	"github.com/DoNewsCode/core/config"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/DoNewsCode/core/srvhttp"
)

// Register the global options includes modules, module constructors and global dependencies
func Register() []Option {
	return []Option{
		/* Dependencies */
		Dependencies(
			app.Providers(),
			otgorm.Providers(),
		),

		/* Module Constructors */
		Constructors(
			docs.New,            // docs
			config.New,          // config module
			core.NewServeModule, // server module
			otgorm.New,
			app.New,
		),

		/* Modules */
		Modules(
			srvhttp.HealthCheckModule{}, // health check module (http demo)
		),
	}
}
