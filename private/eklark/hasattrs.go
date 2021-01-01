//Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

var (
	hasAttrsErrorClass = ekerrors.MustRegisterClass("HasAttrsGetError")
)

func HasAttrsGetString(ha starlark.HasAttrs, attrname string) (starlark.String, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return starlark.String(""), err
	}
	str, ok := value.(starlark.String)
	if !ok {
		return starlark.String(""), ekerrors.New(
			hasAttrsErrorClass,
			attrname+" is not of type string",
		)
	}
	return str, nil
}

func HasAttrsMustGetString(ha starlark.HasAttrs, attrname string) starlark.String {
	str, err := HasAttrsGetString(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return str
}

func HasAttrsGetInt(ha starlark.HasAttrs, attrname string) (starlark.Int, error) {
	value, err := ha.Attr(attrname)
	if err != nil {
		return starlark.MakeInt(0), err
	}
	n, ok := value.(starlark.Int)
	if !ok {
		return starlark.MakeInt(0), ekerrors.New(
			hasAttrsErrorClass,
			attrname+" is not of type int",
		)
	}
	return n, nil
}

func HasAttrsMustGetInt(ha starlark.HasAttrs, attrname string) starlark.Int {
	n, err := HasAttrsGetInt(ha, attrname)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func HasAttrsMustGetGoInt(ha starlark.HasAttrs, attrname string) int {
	nvalue := HasAttrsMustGetInt(ha, attrname)
	n, err := starlark.AsInt32(nvalue)
	if err != nil {
		panic(err.Error())
	}
	return n
}
