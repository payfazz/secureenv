/*
Package secureenv provide utility to access env once.

Motivation

We need to unset the env when it is used, this will prevent environment variable
leak to child process.
*/
package secureenv

import (
	"os"
)

// Get env based on key, if the env not set then return def
func Get(key, def string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	os.Unsetenv(key)
	return env
}
