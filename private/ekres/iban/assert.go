// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertIBAN(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*IBAN)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an IBAN", v)
	}
	return v, nil
}
