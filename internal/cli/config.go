package cli

import (
	"fmt"
	"os"
	"platform/internal/app"
	"platform/internal/infrastructure/configs"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var configType string

var showConfigCmd = &cobra.Command{
	Use:   "configs",
	Short: "Display resolved system configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fxApp := fx.New(
			app.NewCliAppOptions(),

			fx.Invoke(func(cfg *configs.Config) {
				// 1. Convert the entire typed Config struct into a generic map
				rawYaml, err := yaml.Marshal(cfg)
				if err != nil {
					fmt.Printf("Error marshaling base config: %v\n", err)
					return
				}

				var configMap map[string]any
				if err := yaml.Unmarshal(rawYaml, &configMap); err != nil {
					fmt.Printf("Error parsing config to generic map: %v\n", err)
					return
				}

				var dataToMarshal any

				// 2. Dynamically check the requested section without any hardcoded switch cases
				if configType == "" {
					fmt.Println("=== Current Application Configuration ===")
					dataToMarshal = configMap
				} else {
					sectionData, exists := configMap[configType]
					if !exists {
						// Dynamically extract all available keys to show them in the error message
						availableSections := make([]string, 0, len(configMap))
						for k := range configMap {
							availableSections = append(availableSections, k)
						}
						fmt.Printf("Error: unknown config section '%s' (available sections: %v)\n", configType, availableSections)
						os.Exit(1)
					}

					fmt.Printf("=== Current Application Configuration [Section: %s] ===\n", configType)
					dataToMarshal = sectionData
				}

				// 3. Print the final filtered YAML
				finalYaml, err := yaml.Marshal(dataToMarshal)
				if err != nil {
					fmt.Printf("Error marshaling filtered configuration: %v\n", err)
					return
				}

				fmt.Println(string(finalYaml))
			}),
		)

		if err := fxApp.Start(cmd.Context()); err != nil {
			fmt.Printf("Application bootstrap failed: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	showConfigCmd.Flags().StringVarP(
		&configType,
		"section",
		"s",
		"",
		"Filter config output by specific section",
	)

	RootCmd.AddCommand(showConfigCmd)
}
