package bootstrap

import (
	"github.com/DoNewsCode/core"
	"github.com/DoNewsCode/core-starter/internal/cmd"
	"github.com/DoNewsCode/core-starter/internal/config"
	"github.com/spf13/cobra"
	"os"
)

// Bootstrap Project initiated
func Bootstrap() (*cobra.Command, func()) {
	var cfg string

	// Get root command
	root := cmd.NewRootCmd()

	// Determine config path from commandline
	root.PersistentFlags().StringVar(&cfg, "config", "./config.yaml", "config file")
	_ = root.PersistentFlags().Parse(os.Args[1:])

	// Setup core with config path
	c := core.Default(core.WithYamlFile(cfg))

	// Setup global dependencies and register modules
	for _, option := range config.Register() {
		option(c)
	}

	// Apply root command and register commands from modules
	c.ApplyRootCommand(root)

	return root, func() {
		c.Shutdown()
	}
}
