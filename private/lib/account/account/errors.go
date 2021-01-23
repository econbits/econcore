// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("AccountError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*account.Account", "account")
}
