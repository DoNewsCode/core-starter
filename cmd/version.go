package cmd

import (
	"fmt"
	"github.com/DoNewsCode/core"
	"github.com/spf13/cobra"
)

func NewVersionCmd(c *core.C) *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show the version number of this app",
		Long:  `Under semantic versioning, version numbers and the way they change convey meaning about the underlying code and what has been modified from one version to the next.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`
 _____        _   _                   
|  __ \      | \ | |                  
| |  | | ___ |  \| | _____      _____ 
| |  | |/ _ \| . . |/ _ \ \ /\ / / __|
| |__| | (_) | |\  |  __/\ V  V /\__ \
|_____/ \___/|_| \_|\___| \_/\_/ |___/ v%s

%s
`, c.String("version"), c.String("name"))
		},
	}
	return versionCmd
}
