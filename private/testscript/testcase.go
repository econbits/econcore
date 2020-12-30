// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"fmt"
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
)

type TestCase struct {
	Name               string
	FilePath           string
	ExpectedOK         bool
	ExpectedErrorClass *ekerrors.Class
	AbortError         error
}

var (
	maskedErrorClass = ekerrors.MustRegisterClass("Masked Error")
	parseErrorClass  = ekerrors.MustRegisterClass("Parsing Error")
)

func ParseTestCase(fpath string) *TestCase {
	name := eklark.ScriptId(fpath)
	strs := strings.SplitN(name, "_", 2)

	var abortErr error = nil
	var errorClass *ekerrors.Class
	errstring := strs[0]

	if len(strs) != 2 {
		abortErr = ekerrors.New(
			parseErrorClass,
			fmt.Sprintf(
				"filename '%s' does not follow convention: [ErrorType|OK]_case_name.ekm",
				name,
			),
		)
	} else if errstring != "OK" {
		errorClass = ekerrors.MustGetClass(errstring)
	}

	return &TestCase{
		Name:               name,
		FilePath:           fpath,
		ExpectedOK:         errstring == "OK",
		ExpectedErrorClass: errorClass,
		AbortError:         abortErr,
	}
}
