// Copyright (C) 2021  Germán Fuentes Capella

package country

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func HasAttrsGetCountry(ha starlark.HasAttrs, attrname string) (*Country, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return nil, err
	}
	c, ok := value.(*Country)
	if !ok {
		return nil, ekerrors.New(
			errorClass,
			attrname+" is not of type country",
		)
	}
	return c, nil
}

func HasAttrsMustGetCountry(ha starlark.HasAttrs, attrname string) *Country {
	c, err := HasAttrsGetCountry(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return c
}
