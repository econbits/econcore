// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertAccount(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*Account)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an Account", v)
	}
	return v, nil
}
