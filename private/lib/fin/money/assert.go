// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertMoney(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*Money)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an IBAN", v)
	}
	return v, nil
}
