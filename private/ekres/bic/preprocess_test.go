// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestPreProcess(t *testing.T) {
	var value starlark.Value

	value = starlark.String("DEUTDEFFXXX")
	newvalue, err := preprocess(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v, got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = preprocess(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}

	value = starlark.MakeInt(1)
	_, err = preprocess(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
