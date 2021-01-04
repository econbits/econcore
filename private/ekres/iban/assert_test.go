// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertIBAN(t *testing.T) {
	var value starlark.Value
	ibanstr := SampleDE
	value, err := Parse(ibanstr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	newvalue, err := AssertIBAN(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v; got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = AssertIBAN(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
