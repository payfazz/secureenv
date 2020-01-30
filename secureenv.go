/*
Package secureenv provide utility to access env once.

Every function in this package can only be called once, because it
will unset the environment variable.

Every function in this package won't touch output if env is not present.

Motivation

We need to unset the env when it is used, this will prevent environment variable
leak to child process.
*/
package secureenv

import (
	"os"
	"strconv"
	"unsafe"
)

var (
	intsize  = int(unsafe.Sizeof(int(0))) * 8
	uintsize = int(unsafe.Sizeof(uint(0))) * 8
)

// String set data pointed by output.
func String(output *string, key string) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return
	}
	defer os.Unsetenv(key)

	*output = env
}

// Bool set data pointed by output.
//
// return error of strconv.ParseBool.
func Bool(output *bool, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseBool(env)
	if err != nil {
		return err
	}
	*output = val
	return nil
}

// Float64 set data pointed by output.
//
// return error of strconv.ParseFloat.
func Float64(output *float64, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseFloat(env, 64)
	if err != nil {
		return err
	}
	*output = val
	return nil
}

// Float32 set data pointed by output.
//
// return error of strconv.ParseFloat.
func Float32(output *float32, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseFloat(env, 32)
	if err != nil {
		return err
	}
	*output = float32(val)
	return nil
}

// Int64 set data pointed by output.
//
// return error of strconv.ParseInt.
func Int64(output *int64, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseInt(env, 0, 64)
	if err != nil {
		return err
	}
	*output = val
	return nil
}

// Int32 set data pointed by output.
//
// return error of strconv.ParseInt.
func Int32(output *int32, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseInt(env, 0, 32)
	if err != nil {
		return err
	}
	*output = int32(val)
	return nil
}

// Int set data pointed by output.
//
// return error of strconv.ParseInt.
func Int(output *int, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseInt(env, 0, intsize)
	if err != nil {
		return err
	}
	*output = int(val)
	return nil
}

// Uint64 set data pointed by output.
//
// return error of strconv.ParseInt.
func Uint64(output *uint64, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseUint(env, 0, 64)
	if err != nil {
		return err
	}
	*output = val
	return nil
}

// Uint32 set data pointed by output.
//
// return error of strconv.ParseInt.
func Uint32(output *uint32, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseUint(env, 0, 32)
	if err != nil {
		return err
	}
	*output = uint32(val)
	return nil
}

// Uint set data pointed by output.
//
// return error of strconv.ParseInt.
func Uint(output *uint, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseUint(env, 0, uintsize)
	if err != nil {
		return err
	}
	*output = uint(val)
	return nil
}
