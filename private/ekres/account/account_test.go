// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"

	"github.com/econbits/econkit/private/ekres/iban"
	"github.com/econbits/econkit/private/lib/iso/bic"
)

func TestIbanAccount(t *testing.T) {
	_iban := iban.MustParse(iban.SampleDE)
	_bic := bic.MustParse(bic.SampleDE)
	name := "test account"
	kind := "checking"
	acc, err := NewIbanAccount(_iban, name, kind, _bic)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if acc.Provider().String() != _bic.String() {
		t.Fatalf("expected %s; found %v", _bic.String(), acc.Provider())
	}

	if acc.Id().String() != _iban.String() {
		t.Fatalf("expected %s; found %v", _iban.String(), acc.Id())
	}

	if acc.Name() != name {
		t.Fatalf("expected %s; found %v", name, acc.Name())
	}

	if acc.Kind() != kind {
		t.Fatalf("expected %s; found %v", kind, acc.Kind())
	}
}

func TestWalletAccount(t *testing.T) {
	id := "id"
	provider := "provider"
	name := "test account"
	acc := NewWalletAccount(id, name, provider)

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

func TestEqual(t *testing.T) {
	_iban := iban.MustParse(iban.SampleDE)
	_bic := bic.MustParse(bic.SampleDE)
	iacc, err := NewIbanAccount(_iban, "name1", "checking", _bic)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	wacc := NewWalletAccount("id", "name2", "provider")

	if !iacc.Equal(iacc) {
		t.Fatalf("%v is not equal to itself", iacc)
	}

	if iacc.Equal(wacc) {
		t.Fatalf("%v is equal to %v", iacc, wacc)
	}
}
