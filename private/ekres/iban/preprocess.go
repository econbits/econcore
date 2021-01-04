// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/country"
	"go.starlark.net/starlark"
)

var (
	ibanRE = regexp.MustCompile(`[A-Z]{2}[0-9]{2}[0-9A-Z]{4,30}`)
)

func preprocess(v starlark.Value) (starlark.Value, error) {
	v, err := eklark.AssertString(v)
	if err != nil {
		return nil, err
	}

	ibanstr, _ := starlark.AsString(v)
	ibanstr = strings.ToUpper(ibanstr)
	ibanstr = strings.ReplaceAll(ibanstr, " ", "")
	if !ibanRE.MatchString(ibanstr) {
		return nil, ekerrors.New(
			errorClass,
			fmt.Sprintf("'%s' is not a valid IBAN", ibanstr),
		)
	}
	_, err = country.Get(ibanstr[0:2])
	if err != nil {
		return nil, err
	}
	return starlark.String(ibanstr), nil
}
