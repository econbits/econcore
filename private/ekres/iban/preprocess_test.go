// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestPreProcess(t *testing.T) {
	var value starlark.Value

	value = starlark.String(SampleDE)
	newvalue, err := preprocess(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	expectIban := strings.ReplaceAll(SampleDE, " ", "")
	gotIban, _ := starlark.AsString(newvalue)
	if gotIban != expectIban {
		t.Fatalf("expected %s, got %s", expectIban, gotIban)
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
