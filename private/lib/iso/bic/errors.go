// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("BICError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*bic.BIC", "bic")
}
