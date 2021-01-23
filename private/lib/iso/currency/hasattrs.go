// Copyright (C) 2021  Germán Fuentes Capella

package currency

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetCurrency(ha starlark.HasAttrs, attrname string) (*Currency, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	c, ok := value.(*Currency)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type currency",
		)
	}
	return c, nil
}

func HasAttrsMustGetCurrency(ha starlark.HasAttrs, attrname string) *Currency {
	c, err := HasAttrsGetCurrency(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return c
}
