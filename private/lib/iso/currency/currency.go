// Copyright (C) 2020  Germán Fuentes Capella

package currency

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type Currency struct {
	slang.EKValue
}

const (
	currencyType   = "Currency"
	fCurrencyName  = "name"
	fCurrencyCode  = "code"
	fCurrencyId    = "id"
	fCurrencyUnits = "units"
	fnName         = "currency"
)

var (
	currencyErrorClass = ekerrors.MustRegisterClass("CurrencyError")
	CurrencyFn         = &slang.Fn{
		Name:     fnName,
		Callback: currencyFn,
	}

	//go:generate go run ../../tools/synciso4217/main.go
	currencies = map[string]*Currency{}
)

func new_(id int, code string, name string, units int) *Currency {
	return &Currency{
		slang.NewEKValue(
			currencyType,
			[]string{
				fCurrencyName,
				fCurrencyCode,
				fCurrencyId,
				fCurrencyUnits,
			},
			map[string]starlark.Value{
				fCurrencyName:  starlark.String(name),
				fCurrencyCode:  starlark.String(code),
				fCurrencyId:    starlark.MakeInt(id),
				fCurrencyUnits: starlark.MakeInt(units),
			},
			map[string]slang.PreProcessFn{
				fCurrencyName:  slang.AssertString,
				fCurrencyCode:  slang.AssertString,
				fCurrencyId:    slang.AssertInt32,
				fCurrencyUnits: slang.AssertInt32,
			},
			slang.NoMaskFn,
		),
	}
}

func Get(code string) (*Currency, error) {
	currency, exists := currencies[code]
	if !exists {
		return nil, ekerrors.New(
			currencyErrorClass,
			fmt.Sprintf("Currency with code %s not found", code),
		)
	}
	return currency, nil
}

func MustGet(code string) *Currency {
	currency, err := Get(code)
	if err != nil {
		panic(err.Error())
	}
	return currency
}

func currencyFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var code starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fCurrencyCode, &code)
	if err != nil {
		return nil, ekerrors.Wrap(
			currencyErrorClass,
			err.Error(),
			err,
		)
	}
	return Get(string(code))
}

func (c *Currency) Name() string {
	str := slang.HasAttrsMustGetString(c, fCurrencyName)
	return string(str)
}

func (c *Currency) Code() string {
	str := slang.HasAttrsMustGetString(c, fCurrencyCode)
	return string(str)
}

func (c *Currency) Id() int {
	return slang.HasAttrsMustGetGoInt(c, fCurrencyId)
}

func (c *Currency) Units() int {
	return slang.HasAttrsMustGetGoInt(c, fCurrencyUnits)
}

func (c *Currency) Equal(oc *Currency) bool {
	return c == oc
}
