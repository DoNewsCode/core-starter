package bootstrap

import (
	"math/rand"
	"time"

	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core-starter/config"
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
	for _, option := range config.Register() {
		option(c)
	}

	// register commands from modules
	c.ApplyRootCommand(rootCmd.Command)

	return rootCmd.Command, func() {
		c.Shutdown()
	}
}
