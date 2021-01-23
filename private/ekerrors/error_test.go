// Copyright (C) 2020  Germán Fuentes Capella

package ekerrors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	sclass, msg := "ERROR", "msg"
	class := MustRegisterClass(sclass)
	defer delete(registry, sclass)

	err := New(class, msg)

	errstr := "[ERROR] msg"
	if err.Error() != errstr {
		t.Fatalf("Expected '%s'; got '%s'", errstr, err.Error())
	}
}

func TestWrapped(t *testing.T) {
	sclass, msg := "ERROR", "msg"
	class := MustRegisterClass(sclass)
	defer delete(registry, sclass)

	format := func(msg string) string {
		return "changed"
	}

	werr := errors.New(msg)
	err := Wrap(class, werr, []Format{format})

	if err.Unwrap() != werr {
		t.Fatalf("Expected '%v'; got '%v'", werr, err.Unwrap())
	}

	errstr := "[ERROR] changed"
	if err.Error() != errstr {
		t.Fatalf("Expected '%s'; got '%s'", errstr, err.Error())
	}
}

func TestSameClass(t *testing.T) {
	sclass, msg := "ERROR", "msg1"

	class := MustRegisterClass(sclass)
	defer delete(registry, sclass)

	err := New(class, msg)

	if !err.HasClass(class) {
		t.Fatal("Expected true; got false")
	}
}

func TestDifferentClass(t *testing.T) {
	sclass1, sclass2, msg := "ERROR1", "ERROR2", "msg1"

	class1 := MustRegisterClass(sclass1)
	defer delete(registry, sclass1)

	class2 := MustRegisterClass(sclass2)
	defer delete(registry, sclass2)

	err := New(class1, msg)

	if err.HasClass(class2) {
		t.Fatal("Expected false; got true")
	}
}
