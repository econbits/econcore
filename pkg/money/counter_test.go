//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"testing"
)

func TestZeroEUR(t *testing.T) {
	m, err := Zero("EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	c := m.Counter()
	if !c.IsPositive() {
		t.Errorf("Zero is marked as negative: %s", c.String())
	}
	if c.IsNegative() {
		t.Errorf("Zero is marked as negative: %s", c.String())
	}
	if c.String() != "0.00EUR" {
		t.Errorf("Expected: 0.00EUR; got: %s", c.String())
	}
}

func TestOneEUR(t *testing.T) {
	m, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	c := m.Counter()
	if !c.IsPositive() {
		t.Errorf("One is marked as negative: %s", c.String())
	}
	if c.IsNegative() {
		t.Errorf("One is marked as negative: %s", c.String())
	}
	if c.String() != "+1.00EUR" {
		t.Errorf("Expected: +1.00EUR; got: %s", c.String())
	}
}

func TestAddAndSub(t *testing.T) {
	mzero, err := Zero("EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	mone, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	c := mzero.Counter()

	// 0 + 1 = 1
	c, err = c.Add(mone)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.String() != "+1.00EUR" {
		t.Errorf("Expected: +1.00EUR; got: %s", c.String())
	}

	// 1 - 1 = 0
	c, err = c.Sub(mone)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.String() != "0.00EUR" {
		t.Errorf("Expected: 0.00EUR; got: %s", c.String())
	}

	// 0 - 1 = -1
	c, err = c.Sub(mone)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if c.String() != "-1.00EUR" {
		t.Errorf("Expected: -1.00EUR; got: %s", c.String())
	}
	if c.IsPositive() {
		t.Errorf("-1 is marked as positive: %s", c.String())
	}
	if !c.IsNegative() {
		t.Errorf("-1 is marked as positive: %s", c.String())
	}

	// -1 - 1 = -2
	c, err = c.Sub(mone)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.String() != "-2.00EUR" {
		t.Errorf("Expected: -2.00EUR; got: %s", c.String())
	}
	if c.IsPositive() {
		t.Errorf("-2 is marked as positive: %s", c.String())
	}
	if !c.IsNegative() {
		t.Errorf("-2 is marked as positive: %s", c.String())
	}

	// -2 + 1 = -1
	c, err = c.Add(mone)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.String() != "-1.00EUR" {
		t.Errorf("Expected: -1.00EUR; got: %s", c.String())
	}

	// -1 + 1 = 0
	c, err = c.Add(mone)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.String() != "0.00EUR" {
		t.Errorf("Expected: 0.00EUR; got: %s", c.String())
	}
	if !c.IsPositive() {
		t.Errorf("Zero is marked as negative: %s", c.String())
	}
}

func TestDifferentCurrency(t *testing.T) {
	mzero, err := Zero("EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	mone, err := New(100, "USD")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	c := mzero.Counter()

	_, err = c.Add(mone)
	if err == nil {
		t.Errorf("Expected error, found none")
	} else {
		cErr, ok := err.(CurrencyMismatchError)
		if !ok {
			t.Errorf("Expected error type: CurrencyMismatchError")
		}
		if len(cErr.Error()) == 0 {
			t.Errorf("Unexpected empty error message")
		}
	}

	_, err = c.Sub(mone)
	if err == nil {
		t.Errorf("Expected error, found none")
	} else {
		cErr, ok := err.(CurrencyMismatchError)
		if !ok {
			t.Errorf("Expected error type: CurrencyMismatchError")
		}
		if len(cErr.Error()) == 0 {
			t.Errorf("Unexpected empty error message")
		}
	}
}

func TestMaxAmount(t *testing.T) {
	maxUint64 := ^uint64(0)
	mmax, err := New(maxUint64, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	mone, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	cmax := mmax.Counter()

	_, err = cmax.Add(mone)
	if err == nil {
		t.Errorf("Expected error, found none")
	} else {
		cErr, ok := err.(AmountOverflowError)
		if !ok {
			t.Errorf("Expected error type: AmountOverflowError")
		}
		if len(cErr.Error()) == 0 {
			t.Errorf("Unexpected empty error message")
		}
	}
}

func TestMinAmount(t *testing.T) {
	mzero, err := Zero("EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	c := mzero.Counter()

	maxUint64 := ^uint64(0)
	mmax, err := New(maxUint64, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	c, err = c.Sub(mmax)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	mone, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	_, err = c.Sub(mone)
	if err == nil {
		t.Errorf("Expected error, found none")
	} else {
		cErr, ok := err.(AmountOverflowError)
		if !ok {
			t.Errorf("Expected error type: AmountOverflowError")
		}
		if len(cErr.Error()) == 0 {
			t.Errorf("Unexpected empty error message")
		}
	}
}

func TestImmutability(t *testing.T) {
	one, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	cone := one.Counter()
	ctwo, err := cone.Add(one)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if ctwo.Money().IsEqual(one) {
		t.Errorf("[Add] Immutability constraints broken for: %v", one)
	}
	czero, err := cone.Sub(one)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if czero.Money().IsEqual(one) {
		t.Errorf("[Sub] Immutability constraints broken for: %v", one)
	}
}
