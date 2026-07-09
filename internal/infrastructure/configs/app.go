package configs

type AppConfig struct {
	EnvName string `yaml:"env_name"`
	Debug   bool   `yaml:"debug"`
	Name    string `yaml:"name"`
	Level   string `yaml:"level"`
	Test    string `yaml:"example"`
}

func (c *AppConfig) IsDevelop() bool {
	if c.EnvName == "develop" {
		return true
	}
	return false
}
