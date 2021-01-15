// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/ekres/iban"
	"github.com/econbits/econkit/private/lib/iso/bic"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type Account struct {
	slang.EKValue
}

const (
	typeName     = "Account"
	fId          = "id"
	fName        = "name"
	fKind        = "kind"
	fProvider    = "provider"
	fIban        = "iban"
	fBic         = "bic"
	ibanFnName   = "iban_account"
	walletFnName = "wallet_account"
)

var (
	IbanFn = &slang.Fn{
		Name:     ibanFnName,
		Callback: ibanFn,
	}
	WalletFn = &slang.Fn{
		Name:     walletFnName,
		Callback: walletFn,
	}
)

func NewIbanAccount(in *iban.IBAN, name string, kind string, bc *bic.BIC) (*Account, error) {
	return NewIbanAccountFromValues(in, starlark.String(name), starlark.String(kind), bc)
}

func NewIbanAccountFromValues(in *iban.IBAN, name starlark.String, kind starlark.String, bc *bic.BIC) (*Account, error) {
	vkind, err := preprocessIbanKind(kind)
	if err != nil {
		return nil, err
	}
	var vbc starlark.Value
	if bc == nil {
		vbc = starlark.None
	} else {
		vbc = starlark.Value(bc)
	}
	acc := &Account{
		slang.NewEKValue(
			typeName,
			[]string{
				fId,
				fName,
				fKind,
				fProvider,
			},
			map[string]starlark.Value{
				fId:       in,
				fName:     name,
				fKind:     vkind,
				fProvider: vbc,
			},
			map[string]slang.PreProcessFn{
				fId:       iban.AssertIBAN,
				fName:     slang.AssertString,
				fKind:     preprocessIbanKind,
				fProvider: bic.AssertOptionalBIC,
			},
			slang.NoMaskFn,
		),
	}
	err = acc.SetAlias(fIban, fId)
	if err != nil {
		return nil, err
	}
	err = acc.SetAlias(fBic, fProvider)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func NewWalletAccount(id string, name string, provider string) *Account {
	return NewWalletAccountFromValues(
		starlark.String(id),
		starlark.String(name),
		starlark.String(provider),
	)
}

func NewWalletAccountFromValues(
	id starlark.String,
	name starlark.String,
	provider starlark.String,
) *Account {
	return &Account{
		slang.NewEKValue(
			typeName,
			[]string{
				fId,
				fName,
				fKind,
				fProvider,
			},
			map[string]starlark.Value{
				fId:       id,
				fName:     name,
				fKind:     starlark.String(walletType),
				fProvider: provider,
			},
			map[string]slang.PreProcessFn{
				fId:       slang.AssertString,
				fName:     slang.AssertString,
				fKind:     preprocessWalletKind,
				fProvider: slang.AssertString,
			},
			slang.NoMaskFn,
		),
	}
}

func ibanFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var in *iban.IBAN
	var name, kind starlark.String
	var bc *bic.BIC
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		fIban, &in,
		fName, &name,
		fKind, &kind,
		fBic+"?", &bc,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return NewIbanAccountFromValues(in, name, kind, bc)
}

func walletFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var id, provider, name starlark.String
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		fId, &id,
		fName, &name,
		fProvider, &provider,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return NewWalletAccountFromValues(id, name, provider), nil
}

func (acc *Account) Provider() starlark.Value {
	return slang.HasAttrsMustGet(acc, fProvider)
}

func (acc *Account) Id() starlark.Value {
	return slang.HasAttrsMustGet(acc, fId)
}

func (acc *Account) Name() string {
	name := slang.HasAttrsMustGetString(acc, fName)
	return string(name)
}

func (acc *Account) Kind() string {
	kind := slang.HasAttrsMustGetString(acc, fKind)
	return string(kind)
}

func (acc *Account) Equal(oacc *Account) bool {
	return acc == oacc || (acc.Provider() == oacc.Provider() && acc.Id() == oacc.Id())
}
