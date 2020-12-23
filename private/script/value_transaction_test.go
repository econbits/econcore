//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestTransactionAmount(t *testing.T) {
	tx := NewTransaction()

	err := tx.SetField(txAmount, starlark.MakeInt(1))
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	err = tx.SetField(txAmount, starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; found none")
	}
}
