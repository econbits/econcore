// Copyright (C) 2021  Germán Fuentes Capella

package country

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type Country struct {
	slang.EKValue
}

const (
	typeName = "Country"
	fName    = "name"
	fAlpha2  = "alpha2"
	fnName   = "country"
)

var (
	errorClass = ekerrors.MustRegisterClass("CountryError")
	Fn         = &slang.Fn{
		Name:     fnName,
		Callback: countryFn,
	}

	//go:generate go run ../../tools/synciso3166/main.go
	countries = map[string]*Country{}
)

func new_(alpha2, name string) *Country {
	return &Country{
		slang.NewEKValue(
			typeName,
			[]string{fName, fAlpha2},
			map[string]starlark.Value{
				fName:   starlark.String(name),
				fAlpha2: starlark.String(alpha2),
			},
			map[string]slang.PreProcessFn{
				fName:   slang.AssertString,
				fAlpha2: slang.AssertString,
			},
			slang.NoMaskFn,
		),
	}
}

func Get(alpha2 string) (*Country, error) {
	country, exists := countries[alpha2]
	if !exists {
		return nil, ekerrors.New(
			errorClass,
			fmt.Sprintf("Country with alpha2 code %s not found", alpha2),
		)
	}
	return country, nil
}

func MustGet(alpha2 string) *Country {
	country, err := Get(alpha2)
	if err != nil {
		panic(err.Error())
	}
	return country
}

func countryFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var alpha2 starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fAlpha2, &alpha2)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return Get(string(alpha2))
}

func (c *Country) String() string {
	return c.Alpha2()
}

func (c *Country) Alpha2() string {
	alpha2 := slang.HasAttrsMustGetString(c, fAlpha2)
	return string(alpha2)
}

func (c *Country) Name() string {
	name := slang.HasAttrsMustGetString(c, fName)
	return string(name)
}

func (c *Country) Equal(oc *Country) bool {
	return c == oc
}
