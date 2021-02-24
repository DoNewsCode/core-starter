package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd() RootCmd {
	var cfgPath string

	rootCmd := &cobra.Command{
		Use:   "skeleton",
		Short: "A Pragmatic and Opinionated Go Application",
		Long:  `Skeleton provides a starting point to write 12-factor Go Applications.`,
	}

	// Determine config path from commandline
	rootCmd.PersistentFlags().StringVar(
		&cfgPath,
		"config",
		"config.yaml",
		"config file (default is ./config.yaml)",
	)
	_ = rootCmd.PersistentFlags().Parse(os.Args[1:])

	cmd := RootCmd{rootCmd, cfgPath}
	return cmd
}

type RootCmd struct {
	*cobra.Command
	cfgPath string
}

func (c RootCmd) GetCfgPath() string {
	return c.cfgPath
}
