package app

import (
    "github.com/DoNewsCode/core"
    "github.com/DoNewsCode/core/config"
)

type Kernel struct {
}

func Modules(c *core.C) {
    c.AddModuleFunc(InjectKernel)
    c.AddModuleFunc(config.New)
}

func Providers(c *core.C)  {
    // c.Provide(otgorm.Provide)
}
