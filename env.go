package oshelper

import (
	"fmt"
	"os"
)

// GetenvDef retrieves the value of the environment variable named by the 'key'.
// If the variable is not present or has an empty value, 'def' is returned.
func GetenvDef(key, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	return env
}

// GetenvErr retrieves the value of the environment variable named by the 'key'.
// If the variable is not present or has an empty value, an error is returned.
func GetenvErr(key string) (string, error) {
	env, exist := os.LookupEnv(key)
	if !exist {
		return "", fmt.Errorf("environment variable '%s' is not present", key)
	}
	if env == "" {
		return "", fmt.Errorf("environment variable '%s' has an empty value", key)
	}
	return env, nil
}
