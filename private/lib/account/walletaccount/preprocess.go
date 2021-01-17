// Copyright (C) 2021  Germán Fuentes Capella

package walletaccount

import (
	"github.com/econbits/econkit/private/lib/account/account"
	"go.starlark.net/starlark"
)

func matchWalletKind(v starlark.Value) (starlark.Value, error) {
	return account.MatchKind(v, wallet_account_types)
}
