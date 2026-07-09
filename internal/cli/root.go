package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "./app",
	Short: "run app",
	Long:  `start common application`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
