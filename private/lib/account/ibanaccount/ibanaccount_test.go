// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"testing"

	"github.com/econbits/econkit/private/lib/iso/bic"
	"github.com/econbits/econkit/private/lib/iso/iban"
)

func TestIbanAccount(t *testing.T) {
	_iban := iban.MustParse(iban.SampleDE)
	_bic := bic.MustParse(bic.SampleDE)
	name := "test account"
	kind := "checking"
	acc, err := New(_iban, name, kind, _bic)
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

func TestIbanAccountWrongType(t *testing.T) {
	_iban := iban.MustParse(iban.SampleDE)
	_bic := bic.MustParse(bic.SampleDE)
	name := "test account"
	kind := "wallet"
	_, err := New(_iban, name, kind, _bic)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestIbanAccountWithoutBIC(t *testing.T) {
	_iban := iban.MustParse(iban.SampleDE)
	var _bic *bic.BIC
	name := "test account"
	kind := "checking"
	acc, err := New(_iban, name, kind, _bic)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expProvider := "None"
	if acc.Provider().String() != expProvider {
		t.Fatalf("expected %s; found %v", expProvider, acc.Provider())
	}
}
