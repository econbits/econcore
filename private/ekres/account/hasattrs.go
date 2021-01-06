// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetAccount(ha starlark.HasAttrs, attrname string) (*Account, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	a, ok := value.(*Account)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type Account",
		)
	}
	return a, nil
}

func HasAttrsMustGetAccount(ha starlark.HasAttrs, attrname string) *Account {
	a, err := HasAttrsGetAccount(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return a
}
