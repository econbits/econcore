// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"testing"

	"github.com/econbits/econkit/private/eklark"
)

func TestSuccessOnRunOKFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
	Run(testCase, func(fpath string) *eklark.EKError { return nil })
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
}

func TestErrorOnRunOKFile(t *testing.T) {
	testCase := ParseTestCase("OK_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	Run(testCase, func(fpath string) *eklark.EKError {
		return &eklark.EKError{
			FilePath:    testCase.FilePath,
			Function:    "Run",
			ErrorType:   eklark.ErrorType("Test"),
			Description: "Error",
		}
	})

	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorWrongTypeOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("ERROR_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	Run(testCase, func(fpath string) *eklark.EKError {
		return &eklark.EKError{
			FilePath:    testCase.FilePath,
			Function:    "Run",
			ErrorType:   eklark.ErrorType("Test"),
			Description: "Error",
		}
	})

	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}

func TestErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("ERROR_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}

	Run(testCase, func(fpath string) *eklark.EKError {
		return &eklark.EKError{
			FilePath:    testCase.FilePath,
			Function:    "Run",
			ErrorType:   eklark.ErrorType("ERROR"),
			Description: "Error",
		}
	})

	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
}

func TestNoErrorOnRunErrorFile(t *testing.T) {
	testCase := ParseTestCase("ERROR_file.ekm")
	if testCase.GotError != nil {
		t.Fatalf("Unexpected Error %v", testCase.GotError)
	}
	Run(testCase, func(fpath string) *eklark.EKError { return nil })
	if testCase.GotError == nil {
		t.Fatal("Expected Error; none found")
	}
}
