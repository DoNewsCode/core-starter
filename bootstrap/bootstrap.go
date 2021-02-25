package bootstrap

import (
	"github.com/nfangxu/core-skeleton/cmd"
	"math/rand"
	"time"

	"github.com/DoNewsCode/core"
	"github.com/spf13/cobra"
)

func Bootstrap(
	providers func(c *core.C),
	modules func(c *core.C),
	commands ...func(c *core.C) *cobra.Command,
) (*cobra.Command, func()) {
	// setup rand seeder
	rand.Seed(time.Now().UnixNano())

	// init rootCmd
	rootCmd := cmd.NewRootCmd()

	// setup core with config file path
	c := core.Default(core.WithYamlFile(rootCmd.GetCfgPath()))

	// setup global dependencies
	providers(c)

	// setup global modules
	modules(c)

	// register commands from modules
	c.ApplyRootCommand(rootCmd.Command)

	// register global commands
	for _, command := range commands {
		rootCmd.Command.AddCommand(command(c))
	}

	return rootCmd.Command, func() {
		c.Shutdown()
	}
}
