package configs

type AppConfig struct {
	EnvName string `yaml:"env_name"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Commit  string `yaml:"commit"`
	BuildAt string `yaml:"build_at"`
	Debug   bool   `yaml:"debug"`
	Level   string `yaml:"level"`
}

func (c *AppConfig) IsDevelop() bool {
	if c.EnvName == "develop" {
		return true
	}
	return false
}
