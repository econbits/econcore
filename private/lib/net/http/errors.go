// Copyright (C) 2021  Germán Fuentes Capella

package http

import (
	"github.com/econbits/econkit/private/ekerrors"
)

var (
	errorClass = ekerrors.MustRegisterClass("HTTPError")
)
