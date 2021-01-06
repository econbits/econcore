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

	newvalue, err := AssertBIC(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v; got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = AssertBIC(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestAssertOptionalBIC(t *testing.T) {
	bicstr := "DEUTDEFFXXX"
	for _, value := range []starlark.Value{MustParse(bicstr), starlark.None} {
		newvalue, err := AssertOptionalBIC(value)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if newvalue != value {
			t.Fatalf("expected %v; got %v", value, newvalue)
		}
	}
}
