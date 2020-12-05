//Copyright (C) 2020  Germán Fuentes Capella

package tx

import (
	"testing"
	"time"

	"github.com/econbits/econkit/pkg/money"
)

type stringer string

func (s stringer) String() string {
	return string(s)
}

func TestTx(t *testing.T) {
	src := stringer("src")
	dst := stringer("dst")
	tx := New(src, dst, money.MustNew(1, "EUR"), time.Now(), "purpose")
	if tx.Source().String() != src.String() {
		t.Errorf("Expected source: 'src'; found '%s'", tx.Source().String())
	}
	if tx.Destination().String() != dst.String() {
		t.Errorf("Expected destination: 'dst'; found '%s'", tx.Destination().String())
	}
	if tx.Value().String() != "0.01EUR" {
		t.Errorf("Expected value: '0.01EUR'; found '%s'", tx.Value().String())
	}
	if tx.BookingDate() != tx.ValueDate() {
		t.Errorf("Booking and value dates must be equal: '%v' != '%v'", tx.BookingDate(), tx.ValueDate())
	}
	if tx.Purpose() != "purpose" {
		t.Errorf("Expected purpose: 'purpose'; found '%s'", tx.Purpose())
	}
}
