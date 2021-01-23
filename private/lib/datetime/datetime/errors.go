// Copyright (C) 2021  Germán Fuentes Capella

package datetime

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("DateTimeError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*datetime.DateTime", "datetime")
}
