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

// GetString is same with String, but return the data instead
func GetString(key string) (output string, ok bool) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return env, ok
	}
	defer os.Unsetenv(key)
	return env, ok
}

// String set data pointed by output.
func String(output *string, key string) {
	val, ok := GetString(key)
	if ok {
		*output = val
	}
}

// GetBool is same with Bool, but return the data instead
func GetBool(key string) (output bool, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return false, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseBool(env)
	return val, ok, err
}

// Bool set data pointed by output.
//
// return error of strconv.ParseBool.
func Bool(output *bool, key string) error {
	val, ok, err := GetBool(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetFloat64 is same with Float64, but return the data instead
func GetFloat64(key string) (output float64, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseFloat(env, 64)
	return val, ok, err
}

// Float64 set data pointed by output.
//
// return error of strconv.ParseFloat.
func Float64(output *float64, key string) error {
	val, ok, err := GetFloat64(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetFloat32 is same with Float32, but return the data instead
func GetFloat32(key string) (output float32, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseFloat(env, 32)
	return float32(val), ok, err
}

// Float32 set data pointed by output.
//
// return error of strconv.ParseFloat.
func Float32(output *float32, key string) error {
	val, ok, err := GetFloat32(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetInt64 is same with Int64, but return the data instead
func GetInt64(key string) (output int64, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseInt(env, 0, 64)
	return val, ok, err
}

// Int64 set data pointed by output.
//
// return error of strconv.ParseInt.
func Int64(output *int64, key string) error {
	val, ok, err := GetInt64(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetInt32 is same with Int32, but return the data instead
func GetInt32(key string) (output int32, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseInt(env, 0, 32)
	return int32(val), ok, err
}

// Int32 set data pointed by output.
//
// return error of strconv.ParseInt.
func Int32(output *int32, key string) error {
	val, ok, err := GetInt32(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetInt is same with Int, but return the data instead
func GetInt(key string) (output int, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseInt(env, 0, intsize)
	return int(val), ok, err
}

// Int set data pointed by output.
//
// return error of strconv.ParseInt.
func Int(output *int, key string) error {
	val, ok, err := GetInt(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetUint64 is same with Uint64, but return the data instead
func GetUint64(key string) (output uint64, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseUint(env, 0, 64)
	return val, ok, err
}

// Uint64 set data pointed by output.
//
// return error of strconv.ParseInt.
func Uint64(output *uint64, key string) error {
	val, ok, err := GetUint64(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetUint32 is same with Uint32, but return the data instead
func GetUint32(key string) (output uint32, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseUint(env, 0, 32)
	return uint32(val), ok, err
}

// Uint32 set data pointed by output.
//
// return error of strconv.ParseInt.
func Uint32(output *uint32, key string) error {
	val, ok, err := GetUint32(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}

// GetUint is same with Uint, but return the data instead
func GetUint(key string) (output uint, ok bool, err error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return 0, ok, nil
	}
	defer os.Unsetenv(key)

	val, err := strconv.ParseUint(env, 0, uintsize)
	return uint(val), ok, err
}

// Uint set data pointed by output.
//
// return error of strconv.ParseInt.
func Uint(output *uint, key string) error {
	val, ok, err := GetUint(key)
	if err != nil {
		return err
	}
	if ok {
		*output = val
	}
	return nil
}
