//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"strings"
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
		if m.AmountStr() != "1.00" {
			t.Errorf("Expected amount: 1.00; got: %s", m.AmountStr())
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
		if m.AmountStr() != "1" {
			t.Errorf("Expected amount: 1; got: %s", m.AmountStr())
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
		} else {
			cErr, ok := err.(CurrencyNotFoundError)
			if !ok {
				t.Errorf("Expected error type: CurrencyNotFoundError")
			}
			if len(cErr.Error()) == 0 {
				t.Errorf("Unexpected empty error message")
			}
		}
	}
}

func TestNotMoney(t *testing.T) {
	nms := []string{"abc", "EURO"}
	for _, nm := range nms {
		_, err := Parse(nm)
		if err == nil {
			t.Errorf("Expected error; none found")
		} else {
			aErr, ok := err.(AmountNotFoundError)
			if !ok {
				t.Errorf("Expected error type: AmountNotFoundError")
			}
			if len(aErr.Error()) == 0 {
				t.Errorf("Unexpected empty error message")
			}
		}
	}
}

func Test256Units(t *testing.T) {
	u256 := "1." + strings.Repeat("1", 256) + "EUR"
	_, err := Parse(u256)
	if err == nil {
		t.Errorf("Expected error; none found")
	} else {
		aErr, ok := err.(AmountOverflowError)
		if !ok {
			t.Errorf("Expected error type: AmountOverflowError")
		}
		if len(aErr.Error()) == 0 {
			t.Errorf("Unexpected empty error message")
		}
	}
}

func TestEURWith3Units(t *testing.T) {
	_, err := Parse("1.111 EUR")
	if err == nil {
		t.Errorf("Expected error; none found")
	} else {
		tErr, ok := err.(TooManyUnitsError)
		if !ok {
			t.Errorf("Expected error type: TooManyUnitsError")
		}
		if len(tErr.Error()) == 0 {
			t.Errorf("Unexpected empty error message")
		}
	}
}

func TestWrongCurrency(t *testing.T) {
	_, err := Parse("1.11 ABC")
	if err == nil {
		t.Errorf("Expected error; none found")
	} else {
		_, ok := err.(CurrencyNotFoundError)
		if !ok {
			t.Errorf("Expected error type: CurrencyNotFoundError")
		}
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
