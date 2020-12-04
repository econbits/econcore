//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"testing"
)

func TestIBANAccount(t *testing.T) {
	iban := MustParseIBAN(ibanBE)
	bic := MustParseBIC("DEUTDEFF500")
	name := "Test Name"
	typ := TypeGiro
	acc := NewIBANAccount(name, typ, iban, bic)
	if acc.Name() != name {
		t.Errorf("Expected %s; got %s", name, acc.Name())
	}
	if acc.Id().String() != iban.String() {
		t.Errorf("Expected %s; got %s", iban.String(), acc.Id().String())
	}
	if acc.Type() != typ {
		t.Errorf("Expected %d; got %d", typ, acc.Type())
	}
	if acc.Provider().String() != bic.String() {
		t.Errorf("Expected %s; got %s", bic.String(), acc.Provider().String())
	}
}
