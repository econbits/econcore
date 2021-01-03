// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetBIC(ha starlark.HasAttrs, attrname string) (*BIC, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	b, ok := value.(*BIC)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type BIC",
		)
	}
	return b, nil
}

func HasAttrsMustGetBIC(ha starlark.HasAttrs, attrname string) *BIC {
	b, err := HasAttrsGetBIC(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return b
}
