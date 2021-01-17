// Copyright (C) 2021  Germán Fuentes Capella

package walletaccount

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/lib/account/account"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	walletFnName = "wallet_account"
)

var (
	Fn = &slang.Fn{
		Name:     walletFnName,
		Callback: walletFn,
	}
)

func New(id string, name string, provider string) *account.Account {
	return NewFromValues(
		starlark.String(id),
		starlark.String(name),
		starlark.String(provider),
	)
}

func NewFromValues(
	id starlark.String,
	name starlark.String,
	provider starlark.String,
) *account.Account {
	return account.NewFromValues(
		id,
		name,
		starlark.String(account.KindWallet),
		provider,
		slang.AssertString,
		slang.AssertString,
		matchWalletKind,
	)
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
		account.FId, &id,
		account.FName, &name,
		account.FProvider, &provider,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return NewFromValues(id, name, provider), nil
}
