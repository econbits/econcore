// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"fmt"
	"testing"

	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

func TestSuccessOnRunOKFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
	RunTestCase(testCase,
		starlark.StringDict{},
		func(fpath string, epilogue starlark.StringDict) error { return nil },
	)
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
}

func TestErrorOnRunOKFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		func(fpath string, epilogue starlark.StringDict) error {
			return ekerrors.New(
				testscriptErrorClass,
				"Error",
			)
		},
	)

	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestNonEKErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		func(fpath string, epilogue starlark.StringDict) error {
			return fmt.Errorf("this is an error")
		},
	)

	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorWrongTypeOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		func(fpath string, epilogue starlark.StringDict) error {
			return ekerrors.New(
				altTestscriptErrorClass,
				"Error",
			)
		},
	)

	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		func(fpath string, epilogue starlark.StringDict) error {
			return ekerrors.New(
				testscriptErrorClass,
				"Error",
			)
		},
	)

	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
}

func TestNoErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		func(fpath string, epilogue starlark.StringDict) error { return nil },
	)

	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorOnMissingFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
	RunTestCase(testCase,
		starlark.StringDict{},
		ExecScriptFn,
	)
	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}
