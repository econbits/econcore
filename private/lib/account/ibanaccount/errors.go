// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("IBANAccountError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*account.Account", "iban_account")
}
