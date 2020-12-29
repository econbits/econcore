// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"fmt"

	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type TestFn func(path string, epilogue starlark.StringDict) *eklark.EKError

func ExecScriptFn(path string, epilogue starlark.StringDict) *eklark.EKError {
	thread := eklark.NewThread(path)
	_, err := starlark.ExecFile(thread, path, nil, epilogue)
	if err != nil {
		evalerr, ok := err.(*starlark.EvalError)
		if ok {
			ekerr, ok := evalerr.Unwrap().(*eklark.EKError)
			if ok {
				return ekerr
			}
		}
		return &eklark.EKError{
			FilePath:    path,
			Function:    "ExecScriptFn",
			ErrorType:   eklark.ErrorType("TestExecScript"),
			Description: err.Error(),
		}
	}
	return nil
}

func Run(testCase *TestCase, epilogue starlark.StringDict, testFn TestFn) {
	err := testFn(testCase.FilePath, epilogue)
	if testCase.ExpectedOK {
		if err != nil {
			testCase.GotError = err
		}
	} else {
		if err == nil {
			testCase.GotError = &eklark.EKError{
				FilePath:  testCase.FilePath,
				Function:  "Run",
				ErrorType: eklark.ErrorType("Test"),
				Description: fmt.Sprintf(
					"[%s] Expected Error Type %v; none found",
					testCase.FilePath,
					testCase.ExpectedErrorType,
				),
			}
		} else if err.ErrorType != testCase.ExpectedErrorType {
			testCase.GotError = &eklark.EKError{
				FilePath:  testCase.FilePath,
				Function:  "Run",
				ErrorType: eklark.ErrorType("Test"),
				Description: fmt.Sprintf(
					"[%s] Expected Error Type '%v'; found '%v' (Error msg: '%v')",
					testCase.FilePath,
					testCase.ExpectedErrorType,
					err.ErrorType,
					err,
				),
			}
		}
	}
}
