// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetTransaction(ha starlark.HasAttrs, attrname string) (*Transaction, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	o, ok := value.(*Transaction)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type Transaction",
		)
	}
	return o, nil
}

func HasAttrsMustGetTransaction(ha starlark.HasAttrs, attrname string) *Transaction {
	o, err := HasAttrsGetTransaction(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return o
}
