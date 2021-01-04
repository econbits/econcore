// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"fmt"
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/country"
	"go.starlark.net/starlark"
)

func preprocess(v starlark.Value) (starlark.Value, error) {
	v, err := eklark.AssertString(v)
	if err != nil {
		return nil, err
	}

	vstr, _ := starlark.AsString(v)
	code := strings.ToUpper(string(vstr))

	lbic := len(code)
	if lbic != 8 && lbic != 11 {
		return nil, ekerrors.New(
			errorClass,
			fmt.Sprintf("BIC length must be 8 or 11 characters, found %d in %s", lbic, code),
		)
	}
	alpha2 := code[4:6]
	_, err = country.Get(alpha2)
	if err != nil {
		return nil, err
	}

	if lbic == 8 {
		code = code + "XXX"
	}

	return starlark.String(code), nil
}
