package oshelper

import (
	"fmt"
	"os"
)

// GetenvDef retrieves the value of the environment variable named by the 'key'.
// If the value is empty or the variable is not present, 'def' is returned.
func GetenvDef(key, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	return env
}

// MustEnv retrieves the value of the environment variable named by the 'key'.
// If the variable is not present in the environment or the value is empty, an error is returned.
func MustEnv(key string) (string, error) {
	env, exist := os.LookupEnv(key)
	if !exist {
		return "", fmt.Errorf("no environment variable '%s'", key)
	}
	if env == "" {
		return "", fmt.Errorf("empty environment variable '%s'", key)
	}
	return env, nil
}
