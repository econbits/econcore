// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/country"
	"go.starlark.net/starlark"
)

func AssertBIC(v starlark.Value) error {
	_, ok := v.(*BIC)
	if !ok {
		return fmt.Errorf("'%v' is not a BIC", v)
	}
	return nil
}

func assertFormat(code string) error {
	lbic := len(code)
	if lbic != 8 && lbic != 11 {
		return ekerrors.New(
			errorClass,
			fmt.Sprintf("BIC length must be 8 or 11 characters, not %d", lbic),
		)
	}
	alpha2 := code[4:6]
	_, err := country.Get(alpha2)
	if err != nil {
		return err
	}
	return nil
}

func AssertBICString(v starlark.Value) error {
	err := eklark.AssertString(v)
	if err != nil {
		return err
	}
	strv, _ := v.(starlark.String)
	return assertFormat(string(strv))
}
