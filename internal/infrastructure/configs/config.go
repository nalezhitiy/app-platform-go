package configs

import (
	"fmt"
	"os"
	"platform/internal/utils"
	"strings"

	"github.com/drone/envsubst/v2"
	"github.com/goccy/go-yaml"
)

type Config struct {
	App       AppConfig       `yaml:"app"`
	WebServer WebServerConfig `yaml:"webserver"`
	JWT       JWTConfig       `yaml:"jwt"`
}

func NewConfigs() (*Config, error) {
	yamlFiles, err := utils.FindFilesByExtension("configs", ".yml", ".yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to scan config directory: %w", err)
	}

	var fullYamlBuilder strings.Builder
	// Merge all config files into a single text buffer
	for _, filePath := range yamlFiles {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read config file [%s]: %w", filePath, err)
		}
		fullYamlBuilder.Write(content)
		fullYamlBuilder.WriteString("\n")
	}

	// Substitute all environment variables like ${APP_NAME} in the raw YAML text
	processedYaml, err := envsubst.EvalEnv(fullYamlBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("failed to substitute environment variables: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal([]byte(processedYaml), &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal final config YAML: %w", err)
	}

	return &cfg, nil
}
