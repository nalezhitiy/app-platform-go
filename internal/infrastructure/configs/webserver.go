package configs

import "time"

type WebServerConfig struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout TimeoutConfig `yaml:"timeout"`
}

type TimeoutConfig struct {
	Read  time.Duration `yaml:"read"`
	Write time.Duration `yaml:"write"`
	Idle  time.Duration `yaml:"idle"`
}
