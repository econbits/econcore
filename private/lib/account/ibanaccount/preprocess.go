// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"github.com/econbits/econkit/private/lib/account/account"
	"go.starlark.net/starlark"
)

func matchIbanKind(v starlark.Value) (starlark.Value, error) {
	return account.MatchKind(v, iban_account_types)
}
