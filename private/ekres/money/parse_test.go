// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"math/big"
	"testing"
)

func TestParseOneEuro(t *testing.T) {
	ones := []string{"1EUR", "1 EUR", "1.0EUR", "1.0 EUR"}
	for _, one := range ones {
		m, err := Parse(one)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
		if m.Currency().Code() != "EUR" {
			t.Errorf("Expected currency: EUR; got: %s", m.Currency().Code())
		}
		if big.NewInt(100).Cmp(m.Amount()) != 0 {
			t.Errorf("Expected amount: 1.00; got: %v", m.Amount())
		}
		if m.String() != "1.00EUR" {
			t.Errorf("Expected string: 1.00EUR; got: %s", m.String())
		}
	}
}

func TestParseOneJPY(t *testing.T) {
	ones := []string{"1JPY", "1 JPY"}
	for _, one := range ones {
		m, err := Parse(one)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
		if m.Currency().Code() != "JPY" {
			t.Errorf("Expected currency: JPY; got: %s", m.Currency().Code())
		}
		if big.NewInt(1).Cmp(m.Amount()) != 0 {
			t.Errorf("Expected amount: 1; got: %v", m.Amount())
		}
		if m.String() != "1JPY" {
			t.Errorf("Expected string: 1JPY; got: %s", m.String())
		}
	}
}

func TestNoCurrency(t *testing.T) {
	ones := []string{"1", "1.0"}
	for _, one := range ones {
		_, err := Parse(one)
		if err == nil {
			t.Errorf("Expected error; none found")
		}
	}
}

func TestNotMoney(t *testing.T) {
	nms := []string{"abc", "EURO"}
	for _, nm := range nms {
		_, err := Parse(nm)
		if err == nil {
			t.Errorf("Expected error; none found")
		}
	}
}

func TestEURWith3Units(t *testing.T) {
	_, err := Parse("1.111 EUR")
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}

func TestWrongCurrency(t *testing.T) {
	_, err := Parse("1.11 ABC")
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}

func TestMustParseError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParse did not panic")
		}
	}()

	MustParse("1.11 ABC")
}

func TestMustParseOneEuro(t *testing.T) {
	ones := []string{"1EUR", "1 EUR", "1.0EUR", "1.0 EUR"}
	for _, one := range ones {
		m := MustParse(one)
		if m.Currency().Code() != "EUR" {
			t.Errorf("Expected currency: EUR; got: %s", m.Currency().Code())
		}
	}
}
