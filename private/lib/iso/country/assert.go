// Copyright (C) 2021  Germán Fuentes Capella

package country

import (
	"fmt"
	"go.starlark.net/starlark"
)

func AssertCountry(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*Country)
	if !ok {
		return nil, fmt.Errorf("'%v' is not a country", v)
	}
	return v, nil
}
