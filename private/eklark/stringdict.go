// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"

	"go.starlark.net/starlark"
)

type StringDict starlark.StringDict

func (sd StringDict) GetString(key string) (string, error) {
	value, ok := sd[key]
	if !ok {
		return "", fmt.Errorf("key '%s' not in dict", key)
	}
	if value == starlark.None {
		return "", fmt.Errorf("value None for key '%s' can't be converted to string", key)
	}
	str, ok := starlark.AsString(value)
	if !ok {
		return "", fmt.Errorf("value '%v' for key '%s' isn't a string", value, key)
	}
	return str, nil
}
