package configs

import (
	"fmt"
	"platform/internal/utils"
	"time"
)

type JWTConfig struct {
	Secret         utils.SecretString `yaml:"secret"`
	ExpiredAccess  time.Duration      `yaml:"expired_access"`
	ExpiredRefresh time.Duration      `yaml:"expired_refresh"`
}

func (c *JWTConfig) GetSecret() ([]byte, error) {
	if c.Secret == "" {
		return []byte{}, fmt.Errorf("secret is empty")
	}
	return []byte(c.Secret), nil
}
