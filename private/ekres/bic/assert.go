// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertBIC(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*BIC)
	if !ok {
		return nil, fmt.Errorf("'%v' is not a BIC", v)
	}
	return v, nil
}

func AssertOptionalBIC(v starlark.Value) (starlark.Value, error) {
	if v == starlark.None {
		return v, nil
	}
	return AssertBIC(v)
}
