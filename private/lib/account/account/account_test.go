// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAccount(t *testing.T) {
	id := "id"
	provider := "provider"
	kind := "wallet"
	name := "test account"
	acc := NewFromStringValues(
		starlark.String(id),
		starlark.String(name),
		starlark.String(kind),
		starlark.String(provider),
	)

	if acc.Provider().String() != "\""+provider+"\"" {
		t.Fatalf("expected %s; found %v", provider, acc.Provider())
	}

	if acc.Id().String() != "\""+id+"\"" {
		t.Fatalf("expected %s; found %v", id, acc.Id())
	}

	if acc.Name() != name {
		t.Fatalf("expected %s; found %v", name, acc.Name())
	}

	if acc.Kind() != kind {
		t.Fatalf("expected %s; found %v", kind, acc.Kind())
	}
}

func TestEqual(t *testing.T) {
	acc1 := NewFromStringValues(
		starlark.String("id"),
		starlark.String("name"),
		starlark.String("wallet"),
		starlark.String("provider"),
	)
	acc2 := NewFromStringValues(
		starlark.String("id2"),
		starlark.String("name2"),
		starlark.String("wallet"),
		starlark.String("provider2"),
	)

	if !acc1.Equal(acc1) {
		t.Fatalf("%v is not equal to itself", acc1)
	}

	if acc1.Equal(acc2) {
		t.Fatalf("%v is equal to %v", acc1, acc2)
	}
}
