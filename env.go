package env

import (
	"fmt"
	"os"
	"strings"
)

// MustGetEnv tries to return an environment value, if not present, than using defaultValue
func MustGetEnv(prefix string, key string, defaultValue string) string {
	keyWithPrefix := key
	if prefix != "" {
		keyWithPrefix = fmt.Sprintf("%s_%s", strings.ToUpper(prefix), keyWithPrefix)
	}
	value, existed := os.LookupEnv(keyWithPrefix)
	if !existed {
		return defaultValue
	}
	return value
}
