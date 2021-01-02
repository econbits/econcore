// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"math/big"
	"testing"

	"github.com/econbits/econkit/private/ekres/currency"
)

func TestAdd(t *testing.T) {
	one := big.NewInt(1)
	eur := currency.MustGet("EUR")

	m1 := New(one, eur)

	m2, err := m1.Add(m1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if m1.String() != "0.01EUR" {
		t.Errorf("Expected string: 0.01EUR; got: %s", m1.String())
	}
	if m2.String() != "0.02EUR" {
		t.Errorf("Expected string: 0.02EUR; got: %s", m2.String())
	}
}

func TestAddError(t *testing.T) {
	one := big.NewInt(1)
	eur := currency.MustGet("EUR")
	usd := currency.MustGet("USD")

	m1eur := New(one, eur)
	m1usd := New(one, usd)

	_, err := m1eur.Add(m1usd)
	if err == nil {
		t.Errorf("expected error; none found")
	}
}

func TestSub(t *testing.T) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	eur := currency.MustGet("EUR")

	m0 := New(zero, eur)
	m1 := New(one, eur)

	m2, err := m1.Sub(m1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if m1.String() != "0.01EUR" {
		t.Errorf("Expected string: 0.01EUR; got: %s", m1.String())
	}
	if !m2.Equal(m0) {
		t.Errorf("Expected %v; got: %v", m0, m2)
	}
}

func TestSubError(t *testing.T) {
	one := big.NewInt(1)
	eur := currency.MustGet("EUR")
	usd := currency.MustGet("USD")

	m1eur := New(one, eur)
	m1usd := New(one, usd)

	_, err := m1eur.Sub(m1usd)
	if err == nil {
		t.Errorf("expected error; none found")
	}
}
