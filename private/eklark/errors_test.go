// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"
)

func TestErrorType(t *testing.T) {
	expect := "ERROR"
	etype := ErrorType(expect)
	if etype.String() != expect {
		t.Fatalf("Expected Error Type '%s'; got '%s'", expect, etype.String())
	}
}

func TestEKError(t *testing.T) {
	fpath, function, errorType, descr := "filepath", "func", ErrorType("ERROR"), "description"
	eke := EKError{
		FilePath:    fpath,
		Function:    function,
		ErrorType:   errorType,
		Description: descr,
	}
	errstr := "[filepath][func] description"
	if eke.Error() != errstr {
		t.Fatalf("Expected '%s'; got '%s'", errstr, eke.Error())
	}
}
