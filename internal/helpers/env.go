package helpers

import "os"

func MustFromEnv(key string, fallback string) string {
	value, ok := FromEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func FromEnv(key string) (string, bool) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", false
	}
	return value, true
}
