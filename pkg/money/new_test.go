//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"testing"
)

func TestMustZeroError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustZero did not panic")
		}
	}()

	MustZero("ZZZZ")
}

func TestMustZero(t *testing.T) {
	zero := MustZero("EUR")
	if zero.String() != "0.00EUR" {
		t.Errorf("Expected: 0.00EUR; got: %s", zero.String())
	}
}

func TestOneCent(t *testing.T) {
	one := MustNew(1, "EUR")
	if one.String() != "0.01EUR" {
		t.Errorf("Expected: 0.01EUR; got: %s", one.String())
	}
}
