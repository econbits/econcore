// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"

	"go.starlark.net/starlark"
)

type StringDict starlark.StringDict

func (sd StringDict) GetBool(key string) (bool, error) {
	value, ok := sd[key]
	if !ok {
		return false, fmt.Errorf("key '%s' not in dict", key)
	}
	return bool(value.Truth()), nil
}

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

func (sd StringDict) GetStringOr(key string, defaultString string) (string, error) {
	value, ok := sd[key]
	if !ok {
		return "", fmt.Errorf("key '%s' not in dict", key)
	}
	if value == starlark.None {
		return defaultString, nil
	}
	str, ok := starlark.AsString(value)
	if !ok {
		return "", fmt.Errorf("value '%v' for key '%s' isn't a string", value, key)
	}
	return str, nil
}
