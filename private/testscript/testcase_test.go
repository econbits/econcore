// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"testing"
)

func TestOKTestCase(t *testing.T) {
	tc := ParseTestCase("OK_testcase.ekm")
	if !tc.ExpectedOK {
		t.Fatal("test case is expected to be ok; found error")
	}
	if tc.AbortError != nil {
		t.Fatalf("Unexpected error %v", tc.AbortError)
	}
}

func TestErrorTestCase(t *testing.T) {
	tc := ParseTestCase("TestScriptError_testcase.ekm")
	if tc.ExpectedOK {
		t.Fatal("test case is expected to be not ok; found ok")
	}
	if tc.ExpectedErrorClass != testscriptErrorClass {
		t.Fatalf("Expected Error Type 'TestScriptError'; got %v", tc.ExpectedErrorClass)
	}
	if tc.AbortError != nil {
		t.Fatalf("Unexpected error %v", tc.AbortError)
	}
}

func TestInitializationErrorTestCase(t *testing.T) {
	tc := ParseTestCase("testcase.ekm")
	if tc.ExpectedOK {
		t.Fatal("test case is expected to be not ok; found ok")
	}
	if tc.AbortError == nil {
		t.Fatal("Expecting error initializing test case; got none")
	}
}
