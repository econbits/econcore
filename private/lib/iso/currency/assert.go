// Copyright (C) 2021  Germán Fuentes Capella

package currency

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertCurrency(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*Currency)
	if !ok {
		return nil, fmt.Errorf("'%v' is not a currency", v)
	}
	return v, nil
}
