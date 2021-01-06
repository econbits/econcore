// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

func preprocessKind(v starlark.Value, types []string) (starlark.Value, error) {
	v, err := eklark.AssertString(v)
	if err != nil {
		return nil, err
	}

	reqType, _ := starlark.AsString(v)
	for _, atype := range types {
		if atype == reqType {
			return v, nil
		}
	}
	return nil, ekerrors.New(
		errorClass,
		fmt.Sprintf("'%s' is not a valid account type", reqType),
	)
}

func preprocessIbanKind(v starlark.Value) (starlark.Value, error) {
	return preprocessKind(v, iban_account_types)
}

func preprocessWalletKind(v starlark.Value) (starlark.Value, error) {
	return preprocessKind(v, wallet_account_types)
}
