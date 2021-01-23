// Copyright (C) 2021  Germán Fuentes Capella

package currency

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("CurrencyError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*currency.Currency", "currency")
}
