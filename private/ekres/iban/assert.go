// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"fmt"
	"regexp"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/country"
	"go.starlark.net/starlark"
)

func AssertIBAN(v starlark.Value) error {
	_, ok := v.(*IBAN)
	if !ok {
		return fmt.Errorf("'%v' is not an IBAN", v)
	}
	return nil
}

var (
	ibanRE = regexp.MustCompile(`[A-Z]{2}[0-9]{2}[0-9A-Z]{4,30}`)
)

func assertFormat(ibanstr string) error {
	if !ibanRE.MatchString(ibanstr) {
		return ekerrors.New(
			errorClass,
			fmt.Sprintf("'%s' is not a valid IBAN", ibanstr),
		)
	}
	_, err := country.Get(ibanstr[0:2])
	if err != nil {
		return err
	}
	return nil
}

func AssertIBANString(v starlark.Value) error {
	err := eklark.AssertString(v)
	if err != nil {
		return err
	}
	strv, _ := v.(starlark.String)
	return assertFormat(string(strv))
}
