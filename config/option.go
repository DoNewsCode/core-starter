package config

import "github.com/DoNewsCode/core/di"

type Option func(c Core)

type Core interface {
	Provide(deps di.Deps)
	AddModule(modules ...interface{})
	AddModuleFunc(constructor interface{})
}

func Modules(modules ...interface{}) Option {
	return func(c Core) {
		c.AddModule(modules...)
	}
}

func Constructors(constructors ...interface{}) Option {
	return func(c Core) {
		for i := range constructors {
			c.AddModuleFunc(constructors[i])
		}
	}
}

func Dependencies(providers ...di.Deps) Option {
	return func(c Core) {
		for i := range providers {
			c.Provide(providers[i])
		}
	}
}
