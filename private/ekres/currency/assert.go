// Copyright (C) 2021  Germán Fuentes Capella

package currency

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertCurrency(v starlark.Value) error {
	_, ok := v.(*Currency)
	if !ok {
		return fmt.Errorf("'%v' is not a currency", v)
	}
	return nil
}
