// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetIBAN(ha starlark.HasAttrs, attrname string) (*IBAN, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	i, ok := value.(*IBAN)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type IBAN",
		)
	}
	return i, nil
}

func HasAttrsMustGetIBAN(ha starlark.HasAttrs, attrname string) *IBAN {
	i, err := HasAttrsGetIBAN(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return i
}
