// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package clipboard read/write on clipboard
package clipboard

import (
	"strings"
)

// ReadAll read string from clipboard
func ReadAll() (string, error) {
	return readAll()
}

// WriteAll write string to clipboard
func WriteAll(text string) error {
	return writeAll(text)
}

func InitClipboard(environVars []string) {
	envMap := make(map[string]string)
	for _, env := range environVars {
		parts := strings.Split(env, "=")
		envMap[parts[0]] = parts[1]
	}
	initClipboard(envMap)
}

// Unsupported might be set true during clipboard init, to help callers decide
// whether or not to offer clipboard options.
var Unsupported bool
