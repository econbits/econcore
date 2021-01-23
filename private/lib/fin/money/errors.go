// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("MoneyError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*money.Money", "money")
}
