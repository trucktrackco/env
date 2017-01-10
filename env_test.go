package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustGetEnv(t *testing.T) {
	os.Unsetenv("TEST")
	os.Unsetenv("TEST1")

	os.Setenv("TEST", "VALUE")
	value := MustGetEnv("", "TEST", "")
	assert.Equal(t, "VALUE", value, "Expected VALUE value, instead retrieved %s", value)

	value = MustGetEnv("", "TEST1", "DEFAULT")
	assert.Equal(t, "DEFAULT", value, "Expected DEFAULT value, instead retrieved %s", value)

	os.Setenv("DB_TEST", "DB_VALUE")
	value = MustGetEnv("db", "TEST", "DB_DEFAULT")
	assert.Equal(t, "DB_VALUE", value, "Expected DB_VALUE value, instead retrieved %s", value)

	value = MustGetEnv("db", "TEST1", "")
	assert.Equal(t, "", value, "Expected empty value, instead retrieved %s", value)

}
