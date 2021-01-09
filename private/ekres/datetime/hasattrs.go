// Copyright (C) 2021  Germán Fuentes Capella

package datetime

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetDateTime(ha starlark.HasAttrs, attrname string) (*DateTime, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	o, ok := value.(*DateTime)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type DateTime",
		)
	}
	return o, nil
}

func HasAttrsMustGetDateTime(ha starlark.HasAttrs, attrname string) *DateTime {
	o, err := HasAttrsGetDateTime(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return o
}
