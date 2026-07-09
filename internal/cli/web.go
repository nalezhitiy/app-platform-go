package cli

import (
	"platform/internal/app"

	"github.com/spf13/cobra"
)

var routeFlag string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start HTTP Web Server",
	Run: func(cmd *cobra.Command, args []string) {
		application := app.NewWebAdminApp()
		application.Run()
	},
}

func init() {
	// Додаємо прапорець --route (коротка форма -r), дефолтне значення "all"
	webCmd.Flags().StringVarP(&routeFlag, "route", "r", "all", "Which routes to bind (all, public, admin, portal)")
	RootCmd.AddCommand(webCmd)
}
