package app

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core/srvhttp"
)

type Kernel struct {
}

func Modules(c *core.C) {
	c.AddModuleFunc(InjectKernel)
	c.AddModuleFunc(core.NewServeModule)
	// c.AddModuleFunc(config.New)
	// c.AddModuleFunc(otgorm.New)
	c.AddModule(srvhttp.HealthCheckModule{})
}

func Providers(c *core.C) {
	// c.Provide(otgorm.Provide)
}
