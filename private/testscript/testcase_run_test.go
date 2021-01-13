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
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}
	RunTestCase(
		testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		func(fpath string, epilogue starlark.StringDict, load LoadFn) error { return nil },
	)
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}
}

func TestErrorOnRunOKFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		func(fpath string, epilogue starlark.StringDict, load LoadFn) error {
			return ekerrors.New(
				testscriptErrorClass,
				"Error",
			)
		},
	)

	if testCase.AbortError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestNonEKErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		func(fpath string, epilogue starlark.StringDict, load LoadFn) error {
			return fmt.Errorf("this is an error")
		},
	)

	if testCase.AbortError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorWrongTypeOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		func(fpath string, epilogue starlark.StringDict, load LoadFn) error {
			return ekerrors.New(
				altTestscriptErrorClass,
				"Error",
			)
		},
	)

	if testCase.AbortError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		func(fpath string, epilogue starlark.StringDict, load LoadFn) error {
			return ekerrors.New(
				testscriptErrorClass,
				"Error",
			)
		},
	)

	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}
}

func TestNoErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("TestScriptError_file.ekm")
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}

	RunTestCase(testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		func(fpath string, epilogue starlark.StringDict, load LoadFn) error { return nil },
	)

	if testCase.AbortError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorOnMissingFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.AbortError != nil {
		t.Fatalf("Unexpected Error %v", testCase.AbortError)
	}
	RunTestCase(testCase,
		starlark.StringDict{},
		LoadEmptyFn,
		ExecScriptFn,
	)
	if testCase.AbortError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestLoadEmptyFn(t *testing.T) {
	thread := &starlark.Thread{Name: "LoadTest", Load: LoadEmptyFn}
	sd, err := LoadEmptyFn(thread, "package")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if len(sd) != 0 {
		t.Fatalf("expected empty StringDict %v", sd)
	}
}
