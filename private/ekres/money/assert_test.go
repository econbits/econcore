// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"math/big"
	"testing"

	"github.com/econbits/econkit/private/ekres/currency"
	"go.starlark.net/starlark"
)

func TestAssertMoney(t *testing.T) {
	var value starlark.Value
	value = New(big.NewInt(1), currency.MustGet("EUR"))

	newvalue, err := AssertMoney(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v; got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = AssertMoney(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
