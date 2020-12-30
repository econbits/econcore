// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"testing"

	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

func TestIsScript(t *testing.T) {
	if isScript("/tmp", "/tmp") {
		t.Fatal("/tmp is not a script")
	}
	if isScript("/tmp", "filename.txt") {
		t.Fatal("filename.txt is not a script")
	}
	if !isScript("/tmp", "script.ekm") {
		t.Fatal("script.ekm is a script")
	}
}

func TestTestRunner(t *testing.T) {
	err := testRunner(
		"OK_script.ekm",
		starlark.StringDict{},
		func(path string, epilogue starlark.StringDict) error {
			return nil
		})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestTestRunnerOnErrorFile(t *testing.T) {
	err := testRunner(
		"script.ekm",
		starlark.StringDict{},
		func(path string, epilogue starlark.StringDict) error {
			return nil
		})
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestSuccessfulTestRunScript(t *testing.T) {
	TestingRun(
		t,
		"../../test/ekm/vdefault/000_smalltests/testscript",
		starlark.StringDict{},
		ExecScriptFn,
		Fail,
	)
}

func TestErrorTestRunScript(t *testing.T) {
	dpath := "../../test/ekm/vdefault/000_smalltests/testscript/"
	failed := false
	TestingRun(
		t,
		dpath,
		starlark.StringDict{},
		func(path string, epilogue starlark.StringDict) error {
			return &eklark.EKError{
				FilePath:    path,
				Function:    "TestErrorTestRunScript",
				ErrorType:   eklark.ErrorType("Test"),
				Description: "Test Error",
			}
		},
		func(t *testing.T, err error) {
			failed = true
		},
	)
	if !failed {
		t.Fatal("Expected test failure; it succeeded")
	}
}
