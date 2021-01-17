// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"math/big"
	"testing"
	"time"

	"github.com/econbits/econkit/private/lib/account/walletaccount"
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/lib/fin/money"
	"github.com/econbits/econkit/private/lib/iso/currency"
	"go.starlark.net/starlark"
)

func TestTransactionAttributes(t *testing.T) {
	acc := walletaccount.New("id", "name", "provider")
	value := money.New(big.NewInt(100), currency.MustGet("EUR"))
	dt := datetime.NewFromTime(time.Now())
	purpose := "purpose"
	tx := New(acc, acc, value, dt, dt, starlark.String(purpose))
	if !tx.Sender().Equal(acc) {
		t.Fatalf("expected %v; got %v", acc, tx.Sender())
	}
	if !tx.Receiver().Equal(acc) {
		t.Fatalf("expected %v; got %v", acc, tx.Receiver())
	}
	if !tx.Value().Equal(value) {
		t.Fatalf("expected %v; got %v", value, tx.Value())
	}
	if !tx.ValueDate().Equal(dt) {
		t.Fatalf("expected %v; got %v", dt, tx.ValueDate())
	}
	if !tx.BookingDate().Equal(dt) {
		t.Fatalf("expected %v; got %v", dt, tx.BookingDate())
	}
	if tx.Purpose() != purpose {
		t.Fatalf("expected %v; got %v", purpose, tx.Purpose())
	}
}
