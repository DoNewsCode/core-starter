package bootstrap

import (
    "github.com/nfangxu/core-skeleton/cmd"
    "math/rand"
    "time"

    "github.com/DoNewsCode/core"
    "github.com/spf13/cobra"
)

func Bootstrap() (*cobra.Command, *core.C) {
    // setup rand seeder
    rand.Seed(time.Now().UnixNano())

    // init rootCmd
    rootCmd := cmd.NewRootCmd()

    // setup core with config file path
    c := core.Default(core.WithYamlFile(rootCmd.GetCfgPath()))

    // add commands from modules
    c.ApplyRootCommand(rootCmd.Command)

    return rootCmd.Command, c
}
