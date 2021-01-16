// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetMoney(ha starlark.HasAttrs, attrname string) (*Money, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	o, ok := value.(*Money)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type Money",
		)
	}
	return o, nil
}

func HasAttrsMustGetMoney(ha starlark.HasAttrs, attrname string) *Money {
	o, err := HasAttrsGetMoney(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return o
}
