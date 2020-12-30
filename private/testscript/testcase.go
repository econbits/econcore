// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"fmt"
	"strings"

	"github.com/econbits/econkit/private/eklark"
)

type TestCase struct {
	Name              string
	FilePath          string
	ExpectedOK        bool
	ExpectedErrorType eklark.ErrorType
	GotError          error
}

func (tc *TestCase) EKError() *eklark.EKError {
	if tc.GotError == nil {
		return nil
	}
	err, ok := tc.GotError.(*eklark.EKError)
	if !ok {
		return &eklark.EKError{
			FilePath:    tc.FilePath,
			Function:    "EKError",
			ErrorType:   eklark.ErrorType("MaskedEKError"),
			Description: tc.GotError.Error(),
		}
	}
	return err
}

func ParseTestCase(fpath string) *TestCase {
	name := eklark.ScriptId(fpath)
	strs := strings.SplitN(name, "_", 2)

	var gotErr error = nil
	if len(strs) != 2 {
		gotErr = &eklark.EKError{
			FilePath:  fpath,
			Function:  "ParseTestCase",
			ErrorType: eklark.ErrorType("Test"),
			Description: fmt.Sprintf(
				"filename '%s' does not follow convention: [ErrorType|OK]_case_name.ekm",
				name,
			),
		}
	}
	errstring := strs[0]
	return &TestCase{
		Name:              name,
		FilePath:          fpath,
		ExpectedOK:        errstring == "OK",
		ExpectedErrorType: eklark.ErrorType(errstring),
		GotError:          gotErr,
	}
}
