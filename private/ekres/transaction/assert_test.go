// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"math/big"
	"testing"
	"time"

	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/currency"
	"github.com/econbits/econkit/private/ekres/datetime"
	"github.com/econbits/econkit/private/ekres/money"
	"go.starlark.net/starlark"
)

func TestAssertTransaction(t *testing.T) {
	var value starlark.Value
	wallet := account.NewWalletAccount("id", "name", "provider")
	value = New(
		wallet,
		wallet,
		money.New(big.NewInt(1), currency.MustGet("EUR")),
		datetime.NewFromTime(time.Now()),
		nil,
		"purpose",
	)

	newvalue, err := AssertTransaction(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if newvalue != value {
		t.Fatalf("expected %v; got %v", value, newvalue)
	}

	value = starlark.String("")
	_, err = AssertTransaction(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
