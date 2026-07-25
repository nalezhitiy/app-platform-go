package utils

import "encoding/json"

type SecretString string

const Masked = "********"

// Захист від звичайних логів (fmt, zap)
func (s SecretString) String() string {
	if s == "" {
		return ""
	}
	return Masked
}

func (s SecretString) GoString() string {
	if s == "" {
		return ""
	}
	return Masked
}

// Захист від випадкового виведення в JSON API
func (s SecretString) MarshalJSON() ([]byte, error) {
	if s == "" {
		return json.Marshal("")
	}
	return json.Marshal(Masked)
}

// Працює як для github.com/goccy/go-yaml, так і для gopkg.in/yaml.v3
func (s SecretString) MarshalYAML() (any, error) {
	if s == "" {
		return "", nil
	}
	return Masked, nil
}
