package config

import "github.com/DoNewsCode/core/di"

// Option Global register option
type Option func(c Core)

// Core the application core interface
type Core interface {
	Provide(deps di.Deps)
	AddModule(modules ...interface{})
	AddModuleFunc(constructor interface{})
}

// Modules register modules
func Modules(modules ...interface{}) Option {
	return func(c Core) {
		c.AddModule(modules...)
	}
}

// Constructors register module constructors
func Constructors(constructors ...interface{}) Option {
	return func(c Core) {
		for i := range constructors {
			c.AddModuleFunc(constructors[i])
		}
	}
}

// Dependencies register global dependencies
func Dependencies(providers ...di.Deps) Option {
	return func(c Core) {
		for i := range providers {
			c.Provide(providers[i])
		}
	}
}
