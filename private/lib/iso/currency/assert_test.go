// Copyright (C) 2021  Germán Fuentes Capella

package currency

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertCurrency(t *testing.T) {
	var value starlark.Value
	value, err := Get("EUR")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	newvalue, err := AssertCurrency(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v; got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = AssertCurrency(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
