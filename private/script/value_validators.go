// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

type validatorFunc func(starlark.Value) error

func isStringValue(v starlark.Value) error {
	_, ok := v.(starlark.String)
	if !ok {
		return fmt.Errorf("'%v' is not a string", v)
	}
	return nil
}

func isIntValue(v starlark.Value) error {
	_, ok := v.(starlark.Int)
	if !ok {
		return fmt.Errorf("'%v' is not an int", v)
	}
	return nil
}
