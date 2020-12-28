// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"fmt"

	"github.com/econbits/econkit/private/eklark"
)

type TestFn func(path string) *eklark.EKError

func Run(testCase *TestCase, testFn TestFn) {
	err := testFn(testCase.FilePath)
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
