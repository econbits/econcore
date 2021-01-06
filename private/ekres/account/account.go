// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	ekbic "github.com/econbits/econkit/private/ekres/bic"
	ekiban "github.com/econbits/econkit/private/ekres/iban"
	"go.starlark.net/starlark"
)

type Account struct {
	eklark.EKValue
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
	IbanFn = &eklark.Fn{
		Name:     ibanFnName,
		Callback: ibanFn,
	}
	WalletFn = &eklark.Fn{
		Name:     walletFnName,
		Callback: walletFn,
	}
)

func NewIbanAccount(iban *ekiban.IBAN, name string, kind string, bic *ekbic.BIC) (*Account, error) {
	return NewIbanAccountFromValues(iban, starlark.String(name), starlark.String(kind), bic)
}

func NewIbanAccountFromValues(iban *ekiban.IBAN, name starlark.String, kind starlark.String, bic *ekbic.BIC) (*Account, error) {
	vkind, err := preprocessIbanKind(kind)
	if err != nil {
		return nil, err
	}
	var vbic starlark.Value
	if bic == nil {
		vbic = starlark.None
	} else {
		vbic = starlark.Value(bic)
	}
	acc := &Account{
		eklark.NewEKValue(
			typeName,
			[]string{
				fId,
				fName,
				fKind,
				fProvider,
			},
			map[string]starlark.Value{
				fId:       iban,
				fName:     name,
				fKind:     vkind,
				fProvider: vbic,
			},
			map[string]eklark.PreProcessFn{
				fId:       ekiban.AssertIBAN,
				fName:     eklark.AssertString,
				fKind:     preprocessIbanKind,
				fProvider: ekbic.AssertOptionalBIC,
			},
			eklark.NoMaskFn,
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

func NewWalletAccount(id string, name string, provider string) (*Account, error) {
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
) (*Account, error) {
	return &Account{
		eklark.NewEKValue(
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
			map[string]eklark.PreProcessFn{
				fId:       eklark.AssertString,
				fName:     eklark.AssertString,
				fKind:     preprocessWalletKind,
				fProvider: eklark.AssertString,
			},
			eklark.NoMaskFn,
		),
	}, nil
}

func ibanFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var iban *ekiban.IBAN
	var name, kind starlark.String
	var bic *ekbic.BIC
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		fIban, &iban,
		fName, &name,
		fKind, &kind,
		fBic+"?", &bic,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return NewIbanAccountFromValues(iban, name, kind, bic)
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
	return NewWalletAccountFromValues(id, name, provider)
}

func (acc *Account) Provider() starlark.Value {
	return eklark.HasAttrsMustGet(acc, fProvider)
}

func (acc *Account) Id() starlark.Value {
	return eklark.HasAttrsMustGet(acc, fId)
}

func (acc *Account) Name() string {
	name := eklark.HasAttrsMustGetString(acc, fName)
	return string(name)
}

func (acc *Account) Kind() string {
	kind := eklark.HasAttrsMustGetString(acc, fKind)
	return string(kind)
}

func (acc *Account) Equal(oacc *Account) bool {
	return acc == oacc || (acc.Provider() == oacc.Provider() && acc.Id() == oacc.Id())
}
