package bootstrap

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core-starter/internal/cmd"
	"github.com/DoNewsCode/core-starter/internal/config"
	"github.com/spf13/cobra"
)

func Bootstrap() (*cobra.Command, func()) {
	// init rootCmd and get config path
	root, cfg := cmd.NewRootCmd()

	// setup core with config file path
	c := core.Default(core.WithYamlFile(cfg))

	// setup global dependencies
	for _, option := range config.Register() {
		option(c)
	}

	// register commands from modules
	c.ApplyRootCommand(root)

	return root, func() {
		c.Shutdown()
	}
}
