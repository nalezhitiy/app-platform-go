package configs

import "fmt"

type JWTConfig struct {
	secret         string `yaml:"secret"`
	ExpiredAccess  int    `yaml:"expired_access"`
	ExpiredRefresh int    `yaml:"expired_refresh"`
}

func (c *JWTConfig) GetSecret() (string, error) {
	if c.secret == "" {
		return "", fmt.Errorf("secret is empty")
	}
	return c.secret, nil
}
