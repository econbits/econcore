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

	err = AssertIBAN(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	value = starlark.String("")
	err = AssertIBAN(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestAssertIBANString(t *testing.T) {
	var value starlark.Value

	value = format(starlark.String(SampleDE))
	err := AssertIBANString(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	value = starlark.String("")
	err = AssertIBANString(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}

	value = starlark.MakeInt(1)
	err = AssertIBANString(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
