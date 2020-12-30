// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"fmt"

	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type TestFn func(path string, epilogue starlark.StringDict) error

func ExecScriptFn(path string, epilogue starlark.StringDict) error {
	thread := eklark.NewThread(path)
	_, err := eklark.Exec(thread, path, epilogue)
	if err != nil {
		ekerr, ok := err.(*eklark.EKError)
		if ok {
			return ekerr
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

func RunTestCase(testCase *TestCase, epilogue starlark.StringDict, testFn TestFn) {
	err := testFn(testCase.FilePath, epilogue)
	if testCase.ExpectedOK {
		if err != nil {
			testCase.GotError = err
		}
	} else {
		if err == nil {
			testCase.GotError = &eklark.EKError{
				FilePath:  testCase.FilePath,
				Function:  "RunTestCase",
				ErrorType: eklark.ErrorType("Test"),
				Description: fmt.Sprintf(
					"[%s] Expected Error Type %v; none found",
					testCase.FilePath,
					testCase.ExpectedErrorType,
				),
			}
		} else {
			ekerr, ok := err.(*eklark.EKError)
			if !ok {
				testCase.GotError = &eklark.EKError{
					FilePath:  testCase.FilePath,
					Function:  "RunTestCase",
					ErrorType: eklark.ErrorType("Test"),
					Description: fmt.Sprintf(
						"[%s] Expected Error Type %v; found %v",
						testCase.FilePath,
						testCase.ExpectedErrorType,
						err,
					),
				}
			} else if ekerr.ErrorType != testCase.ExpectedErrorType {
				testCase.GotError = &eklark.EKError{
					FilePath:  testCase.FilePath,
					Function:  "RunTestCase",
					ErrorType: eklark.ErrorType("Test"),
					Description: fmt.Sprintf(
						"[%s] Expected Error Type '%v'; found '%v' (Error msg: '%v')",
						testCase.FilePath,
						testCase.ExpectedErrorType,
						ekerr.ErrorType,
						ekerr,
					),
				}
			}
		}
	}
}
