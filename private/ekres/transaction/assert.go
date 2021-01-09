// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertTransaction(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*Transaction)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an Transaction", v)
	}
	return v, nil
}
