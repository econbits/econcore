// Copyright (C) 2021  Germán Fuentes Capella

package assert

import (
	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("AssertionError")
)

func FormatError(msg string) string {
	return msg
}
