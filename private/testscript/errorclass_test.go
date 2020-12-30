// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"github.com/econbits/econkit/private/ekerrors"
)

var (
	testscriptErrorClass    = ekerrors.MustRegisterClass("TestScriptError")
	altTestscriptErrorClass = ekerrors.MustRegisterClass("Alternative TestScriptError")
)
