// Copyright (C) 2020  Germán Fuentes Capella

package ekerrors

import (
	"testing"
)

func TestUnregisteredClass(t *testing.T) {
	s := "Test"

	class := MustRegisterClass(s)
	defer delete(registry, s)

	if class.s != s {
		t.Fatalf("expected %s; found %s", s, class.s)
	}

	if class != MustGetClass(s) {
		t.Fatalf("expected %s; found %s", class, MustGetClass(s))
	}
}

func TestReRegisteredClass(t *testing.T) {
	s := "Test"

	MustRegisterClass(s)
	defer delete(registry, s)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()
	MustRegisterClass(s)
}

func TestMustGetClassPanic(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	s := "Test"

	MustGetClass(s)
}
