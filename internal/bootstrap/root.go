package bootstrap

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root command.
func NewRootCmd(name, version string) *cobra.Command {
	return &cobra.Command{
		Use:     name,
		Short:   "A Pragmatic and Opinionated Go Application",
		Long:    `Core-starter provides a starting point to write 12-factor Go Applications.`,
		Version: version,
	}
}
