// Copyright (C) 2021  Germán Fuentes Capella

package country

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertCountry(t *testing.T) {
	var value starlark.Value
	value, err := Get("DE")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	err = AssertCountry(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	value = starlark.String("")
	err = AssertCountry(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
