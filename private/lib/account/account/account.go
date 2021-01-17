// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type Account struct {
	slang.EKValue
}

const (
	typeName  = "Account"
	FId       = "id"
	FName     = "name"
	FKind     = "kind"
	FProvider = "provider"
)

func NewFromStringValues(
	id starlark.String,
	name starlark.String,
	kind starlark.String,
	provider starlark.String,
) *Account {
	return NewFromValues(
		id,
		name,
		kind,
		provider,
		slang.AssertString,
		slang.AssertString,
		matchAccountKind,
	)
}

func NewFromValues(
	id starlark.Value,
	name starlark.String,
	kind starlark.String,
	provider starlark.Value,
	idPreProcessor slang.PreProcessFn,
	providerPreProcessor slang.PreProcessFn,
	kindPreProcessor slang.PreProcessFn,
) *Account {
	return &Account{
		slang.NewEKValue(
			typeName,
			[]string{
				FId,
				FName,
				FKind,
				FProvider,
			},
			map[string]starlark.Value{
				FId:       id,
				FName:     name,
				FKind:     kind,
				FProvider: provider,
			},
			map[string]slang.PreProcessFn{
				FId:       idPreProcessor,
				FName:     slang.AssertString,
				FKind:     kindPreProcessor,
				FProvider: providerPreProcessor,
			},
			slang.NoMaskFn,
		),
	}
}

func (acc *Account) Provider() starlark.Value {
	return slang.HasAttrsMustGet(acc, FProvider)
}

func (acc *Account) Id() starlark.Value {
	return slang.HasAttrsMustGet(acc, FId)
}

func (acc *Account) Name() string {
	name := slang.HasAttrsMustGetString(acc, FName)
	return string(name)
}

func (acc *Account) Kind() string {
	kind := slang.HasAttrsMustGetString(acc, FKind)
	return string(kind)
}

func (acc *Account) Equal(oacc *Account) bool {
	return acc == oacc || (acc.Provider() == oacc.Provider() && acc.Id() == oacc.Id())
}
