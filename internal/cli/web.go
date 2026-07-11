package cli

import (
	"fmt"
	"os"
	"platform/internal/app"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var apiFlag string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start HTTP Web Server",
	Long: `Start one or more HTTP API modules.
Examples:
  app web
  app web --api all
  app web --api admin
  app web --api portal`,

	Run: func(cmd *cobra.Command, args []string) {

		var application = func() *fx.App {
			switch apiFlag {

			case "all":
				return app.NewWebAllApp()

			case "admin":
				return app.NewWebAdminApp()

			case "portal":
				return app.NewWebPortalApp()

			default:
				fmt.Printf("Unknown API module: %s\n", apiFlag)
				os.Exit(1)
				return nil
			}
		}()

		application.Run()
	},
}

func init() {
	webCmd.Flags().StringVarP(
		&apiFlag,
		"api",
		"a",
		"all",
		"API modules to start (all, admin, portal)",
	)

	RootCmd.AddCommand(webCmd)
}
