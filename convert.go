package rest

import (
	"encoding/base64"
	"strconv"
)

// ToBool tries to parse the given string to a bool value.
func ToBool(value string) (bool, bool) {
	v, err := strconv.ParseBool(value)
	return v, err == nil
}

// ToFloat32 tries to parse the given string to a float32 value.
func ToFloat32(value string) (float32, bool) {
	v, err := strconv.ParseFloat(value, 32)
	return float32(v), err == nil
}

// ToFloat64 tries to parse the given string to a float64 value.
func ToFloat64(value string) (float64, bool) {
	v, err := strconv.ParseFloat(value, 64)
	return float64(v), err == nil
}

// ToString returns the given argument if string not empty.
func ToString(value string) (string, bool) {
	return value, len(value) > 0
}

// ToInt tries to parse the given string to a int value.
func ToInt(value string) (int, bool) {
	v, err := strconv.ParseInt(value, 10, 0)
	return int(v), err == nil
}

// ToInt32 tries to parse the given string to a int32 value.
func ToInt32(value string) (int32, bool) {
	v, err := strconv.ParseInt(value, 10, 32)
	return int32(v), err == nil
}

// ToInt64 tries to parse the given string to a int64 value.
func ToInt64(value string) (int64, bool) {
	v, err := strconv.ParseInt(value, 10, 64)
	return int64(v), err == nil
}

// ToUint tries to parse the given string to a uint value.
func ToUint(value string) (uint, bool) {
	v, err := strconv.ParseUint(value, 10, 0)
	return uint(v), err == nil
}

// ToUint32 tries to parse the given string to a uint32 value.
func ToUint32(value string) (uint32, bool) {
	v, err := strconv.ParseUint(value, 10, 32)
	return uint32(v), err == nil
}

// ToUint64 tries to parse the given string to a uint64 value.
func ToUint64(value string) (uint64, bool) {
	v, err := strconv.ParseUint(value, 10, 64)
	return uint64(v), err == nil
}

// ToBytes tries to parse the given string to a bytes value.
func ToBytes(value string) ([]byte, bool) {
	decoded, err := base64.URLEncoding.DecodeString(value)
	return decoded, err == nil
}
