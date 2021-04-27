package cmd

import (
	"github.com/spf13/cobra"
)

const (
	command = "root"
	version = "v1.0.0"
)

// NewRootCmd Define the root command
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:     command,
		Short:   "A Pragmatic and Opinionated Go Application",
		Long:    `Core-starter provides a starting point to write 12-factor Go Applications.`,
		Version: version,
	}
}
