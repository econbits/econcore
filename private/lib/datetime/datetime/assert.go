// Copyright (C) 2020  Germán Fuentes Capella

package datetime

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertDateTime(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*DateTime)
	if !ok {
		return nil, fmt.Errorf("'%v' is not a date", v)
	}
	return v, nil
}
