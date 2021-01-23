// Copyright (C) 2021  Germán Fuentes Capella

package walletaccount

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("WalletAccountError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*account.Account", "wallet_account")
}
