// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/lib/account/account"
	"github.com/econbits/econkit/private/lib/iso/bic"
	"github.com/econbits/econkit/private/lib/iso/iban"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	fIban      = "iban"
	fBic       = "bic"
	fName      = "name"
	fKind      = "kind"
	ibanFnName = "iban_account"
)

var (
	Fn = &slang.Fn{
		Name:     ibanFnName,
		Callback: ibanFn,
	}
)

func New(in *iban.IBAN, name string, kind string, bc *bic.BIC) (*account.Account, error) {
	return NewFromValues(in, starlark.String(name), starlark.String(kind), bc)
}

func NewFromValues(in *iban.IBAN, name starlark.String, kind starlark.String, bc *bic.BIC) (*account.Account, error) {
	_, err := matchIbanKind(kind)
	if err != nil {
		return nil, err
	}
	var vbc starlark.Value
	if bc == nil {
		vbc = starlark.None
	} else {
		vbc = starlark.Value(bc)
	}
	acc := account.NewFromValues(
		in,
		name,
		kind,
		vbc,
		iban.AssertIBAN,
		bic.AssertOptionalBIC,
		matchIbanKind,
	)
	err = acc.SetAlias(fIban, account.FId)
	if err != nil {
		return nil, err
	}
	err = acc.SetAlias(fBic, account.FProvider)
	if err != nil {
		return nil, err
	}
	return acc, nil
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
	return NewFromValues(in, name, kind, bc)
}
