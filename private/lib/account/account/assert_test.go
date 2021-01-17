// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertAccount(t *testing.T) {
	var value starlark.Value
	value = NewFromStringValues("id", "name", "wallet", "provider")

	newvalue, err := AssertAccount(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v; got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = AssertAccount(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
