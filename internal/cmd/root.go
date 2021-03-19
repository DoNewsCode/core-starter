package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	command = "root"
	version = "v1.0.0"
)

func NewRootCmd() (*cobra.Command, string) {
	var cfgPath string

	rootCmd := &cobra.Command{
		Use:     command,
		Short:   "A Pragmatic and Opinionated Go Application",
		Long:    `Skeleton provides a starting point to write 12-factor Go Applications.`,
		Version: version,
	}

	// Determine config path from commandline
	rootCmd.PersistentFlags().StringVar(
		&cfgPath,
		"config",
		"./config.yaml",
		"config file (default is ./config.yaml)",
	)
	_ = rootCmd.PersistentFlags().Parse(os.Args[1:])

	return rootCmd, cfgPath
}
