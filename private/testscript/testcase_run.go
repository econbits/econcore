// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"errors"
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

var (
	runTCErrorClass = ekerrors.MustRegisterClass("TestScript TestCase")
)

type TestFn func(path string, epilogue starlark.StringDict) error

func ExecScriptFn(path string, epilogue starlark.StringDict) error {
	thread := &starlark.Thread{Name: path}
	_, err := starlark.ExecFile(thread, path, nil, epilogue)
	return err
}

func RunTestCase(testCase *TestCase, epilogue starlark.StringDict, testFn TestFn) {
	err := testFn(testCase.FilePath, epilogue)
	if testCase.ExpectedOK {
		if err != nil {
			testCase.GotError = err
		}
	} else {
		if err == nil {
			testCase.GotError = ekerrors.New(
				runTCErrorClass,
				fmt.Sprintf(
					"[%s] Expected Error Type %v; none found",
					testCase.FilePath,
					testCase.ExpectedErrorType,
				),
			)
		} else {
			var ekerr *ekerrors.EKError
			if errors.As(err, &ekerr) {
				if !ekerr.HasClass(testCase.ExpectedErrorType) {
					testCase.GotError = ekerrors.New(
						runTCErrorClass,
						fmt.Sprintf(
							"[%s] Expected Error '%v'; found '%v')",
							testCase.FilePath,
							testCase.ExpectedErrorType,
							ekerr,
						),
					)
				}
			} else {
				testCase.GotError = ekerrors.Wrap(
					runTCErrorClass,
					fmt.Sprintf(
						"[%s] Expected Error Type %v; found %v",
						testCase.FilePath,
						testCase.ExpectedErrorType,
						err,
					),
					err,
				)
			}
		}
	}
}
