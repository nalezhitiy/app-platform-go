package utils

type SecretString string

var MASKED = "********"

// Захист від звичайних логів (fmt, zap)
func (s SecretString) String() string {
	return MASKED
}

func (s SecretString) GoString() string {
	return MASKED
}

// Захист від випадкового виведення в JSON API
func (s SecretString) MarshalJSON() ([]byte, error) {
	return []byte(MASKED), nil
}

// для github.com/goccy/go-yaml, так і для gopkg.in/yaml.v3
func (s SecretString) MarshalYAML() (any, error) {
	return MASKED, nil
}
