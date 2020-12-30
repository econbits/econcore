// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"testing"

	"github.com/econbits/econkit/private/eklark"
)

func TestOKTestCase(t *testing.T) {
	tc := ParseTestCase("OK_testcase.ekm")
	if !tc.ExpectedOK {
		t.Fatal("test case is expected to be ok; found error")
	}
	if tc.GotError != nil {
		t.Fatalf("Unexpected error %v", tc.GotError)
	}
	if tc.EKError() != nil {
		t.Fatalf("Unexpected error %v", tc.EKError())
	}
}

func TestErrorTestCase(t *testing.T) {
	tc := ParseTestCase("ERROR_testcase.ekm")
	if tc.ExpectedOK {
		t.Fatal("test case is expected to be not ok; found ok")
	}
	if tc.ExpectedErrorType != eklark.ErrorType("ERROR") {
		t.Fatalf("Expected Error Type 'ERROR'; got %v", tc.ExpectedErrorType)
	}
	if tc.GotError != nil {
		t.Fatalf("Unexpected error %v", tc.GotError)
	}
	if tc.EKError() != nil {
		t.Fatalf("Unexpected error %v", tc.EKError())
	}
}

func TestInitializationErrorTestCase(t *testing.T) {
	tc := ParseTestCase("testcase.ekm")
	if tc.ExpectedOK {
		t.Fatal("test case is expected to be not ok; found ok")
	}
	if tc.GotError == nil {
		t.Fatal("Expecting error initializing test case; got none")
	}
	if tc.EKError() == nil {
		t.Fatal("Expecting error initializing test case; got none")
	}
}
