package configs

import (
	"fmt"
	"platform/internal/utils"
)

type JWTConfig struct {
	Secret         utils.SecretString `yaml:"secret"`
	ExpiredAccess  int                `yaml:"expired_access"`
	ExpiredRefresh int                `yaml:"expired_refresh"`
}

func (c *JWTConfig) GetSecret() ([]byte, error) {
	if c.Secret == "" {
		return []byte{}, fmt.Errorf("secret is empty")
	}
	return []byte(c.Secret), nil
}
