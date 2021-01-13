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

type LoadFn func(thread *starlark.Thread, module string) (starlark.StringDict, error)

func LoadEmptyFn(_ *starlark.Thread, _ string) (starlark.StringDict, error) {
	return starlark.StringDict{}, nil
}

type TestFn func(path string, epilogue starlark.StringDict, load LoadFn) error

func ExecScriptFn(path string, epilogue starlark.StringDict, load LoadFn) error {
	thread := &starlark.Thread{Name: path, Load: load}
	_, err := starlark.ExecFile(thread, path, nil, epilogue)
	return err
}

func RunTestCase(testCase *TestCase, epilogue starlark.StringDict, load LoadFn, testFn TestFn) {
	err := testFn(testCase.FilePath, epilogue, load)
	if testCase.ExpectedOK {
		if err != nil {
			testCase.AbortError = err
		}
	} else {
		if err == nil {
			testCase.AbortError = ekerrors.New(
				runTCErrorClass,
				fmt.Sprintf(
					"[%s] Expected Error Type %v; none found",
					testCase.FilePath,
					testCase.ExpectedErrorClass,
				),
			)
		} else {
			var ekerr *ekerrors.EKError
			if errors.As(err, &ekerr) {
				if !ekerr.HasClass(testCase.ExpectedErrorClass) {
					testCase.AbortError = ekerrors.New(
						runTCErrorClass,
						fmt.Sprintf(
							"[%s] Expected Error '%v'; found '%v')",
							testCase.FilePath,
							testCase.ExpectedErrorClass,
							ekerr,
						),
					)
				}
			} else {
				testCase.AbortError = ekerrors.Wrap(
					runTCErrorClass,
					fmt.Sprintf(
						"[%s] Expected Error Type %v; found %v",
						testCase.FilePath,
						testCase.ExpectedErrorClass,
						err,
					),
					err,
				)
			}
		}
	}
}
