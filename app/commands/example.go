package commands

import (
	"fmt"
	"github.com/DoNewsCode/core/contract"
	"github.com/spf13/cobra"
)

// NewExampleCommand Registers a example command
func NewExampleCommand(c contract.ConfigAccessor) *cobra.Command {
	return &cobra.Command{
		Use:   "example",
		Short: "Example command.",
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
}
