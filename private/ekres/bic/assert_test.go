// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertBIC(t *testing.T) {
	var value starlark.Value
	bicstr := "DEUTDEFFXXX"
	value, err := Parse(bicstr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	err = AssertBIC(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	value = starlark.String("")
	err = AssertBIC(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestAssertBICString(t *testing.T) {
	var value starlark.Value

	value = starlark.String("DEUTDEFFXXX")
	err := AssertBICString(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	value = starlark.String("")
	err = AssertBICString(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}

	value = starlark.MakeInt(1)
	err = AssertBICString(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
