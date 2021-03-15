package bootstrap

import (
	"fmt"
	"github.com/DoNewsCode/core-starter/config"
	"math/rand"
	"reflect"
	"time"

	"github.com/DoNewsCode/core"
	"github.com/spf13/cobra"
)

func Bootstrap() (*cobra.Command, func()) {
	// setup rand seeder
	rand.Seed(time.Now().UnixNano())

	// init rootCmd
	rootCmd := NewRootCmd()

	// setup core with config file path
	c := core.Default(core.WithYamlFile(rootCmd.GetCfgPath()))

	// setup global dependencies
	for _, provider := range config.Providers {
		c.Provide(provider)
	}

	// setup global modules
	for _, module := range config.Modules {
		t := reflect.TypeOf(module)

		switch true {
		case t.Kind() == reflect.Func:
			c.AddModuleFunc(module)
		case t.Kind() == reflect.Struct:
			c.AddModule(module)
		default:
			panic(fmt.Sprintf("Unexpected module [%s]", t.Kind()))
		}
	}

	// register commands from modules
	c.ApplyRootCommand(rootCmd.Command)

	return rootCmd.Command, func() {
		c.Shutdown()
	}
}
