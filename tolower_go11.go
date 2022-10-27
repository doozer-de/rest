//go:build !go1.12
// +build !go1.12

package rest

import "strings"

func toLower(s string) string {
	return strings.ToLower(s)
}
