// Copyright (C) 2020  Germán Fuentes Capella

package currency

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type Currency struct {
	eklark.EKValue
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
	CurrencyFn         = &eklark.Fn{
		Name:     fnName,
		Callback: currencyFn,
	}

	//go:generate go run ../../tools/synciso4217/main.go
	currencies = map[string]*Currency{}
)

func new_(id int, code string, name string, units int) *Currency {
	return &Currency{
		eklark.NewEKValue(
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
			map[string]eklark.ValidateFn{
				fCurrencyName:  eklark.AssertString,
				fCurrencyCode:  eklark.AssertString,
				fCurrencyId:    eklark.AssertInt32,
				fCurrencyUnits: eklark.AssertInt32,
			},
			eklark.NoMaskFn,
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
	str := eklark.HasAttrsMustGetString(c, fCurrencyName)
	return string(str)
}

func (c *Currency) Code() string {
	str := eklark.HasAttrsMustGetString(c, fCurrencyCode)
	return string(str)
}

func (c *Currency) Id() int {
	return eklark.HasAttrsMustGetGoInt(c, fCurrencyId)
}

func (c *Currency) Units() int {
	return eklark.HasAttrsMustGetGoInt(c, fCurrencyUnits)
}

func (c *Currency) Equal(oc *Currency) bool {
	return c == oc
}
