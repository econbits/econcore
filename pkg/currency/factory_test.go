//Copyright (C) 2020  Germán Fuentes Capella

package currency

import (
	"testing"
)

func TestWrongCurrencyCodes(t *testing.T) {
	wrongCodes := []string{"", "___"}
	for _, code := range wrongCodes {
		c, err := Get(code)
		if err == nil {
			t.Errorf("Expected Error, got nil")
		}
		if len(c.Name()) > 0 {
			t.Errorf("Expected empty name, got %s", c.Name())
		}
		if len(c.Code()) > 0 {
			t.Errorf("Expected empty currency code, got %s", c.Code())
		}
		if c.Id() > 0 {
			t.Errorf("Expected currency id: 0, got %d", c.Id())
		}
		if c.Units() > 0 {
			t.Errorf("Expected currency units: 0, got %d", c.Units())
		}
	}
}

func TestEuro(t *testing.T) {
	c, err := Get("EUR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.Name() != "Euro" {
		t.Errorf("Expected name: Euro, got %s", c.Name())
	}
	if c.Code() != "EUR" {
		t.Errorf("Expected currency code: EUR, got %s", c.Code())
	}
	if c.Id() != 978 {
		t.Errorf("Expected currency id: 978, got %d", c.Id())
	}
	if c.Units() != 2 {
		t.Errorf("Expected currency units: 2, got %d", c.Units())
	}
}
