// Copyright (C) 2021  Germán Fuentes Capella

package walletaccount

import (
	"testing"
)

func TestWalletAccount(t *testing.T) {
	id := "id"
	provider := "provider"
	name := "test account"
	acc := New(id, name, provider)

	if acc.Provider().String() != "\""+provider+"\"" {
		t.Fatalf("expected %s; found %v", provider, acc.Provider())
	}

	if acc.Id().String() != "\""+id+"\"" {
		t.Fatalf("expected %s; found %v", id, acc.Id())
	}

	if acc.Name() != name {
		t.Fatalf("expected %s; found %v", name, acc.Name())
	}

	kind := "wallet"
	if acc.Kind() != kind {
		t.Fatalf("expected %s; found %v", kind, acc.Kind())
	}
}
