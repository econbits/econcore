// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("IBANError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*iban.IBAN", "iban")
}
