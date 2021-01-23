// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("TransactionError")
)

func FormatError(msg string) string {
	return strings.ReplaceAll(msg, "*transaction.Transaction", "transaction")
}
