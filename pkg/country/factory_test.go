//Copyright (C) 2020  Germán Fuentes Capella

package country

import (
	"testing"
)

func TestWrongCountryCodes(t *testing.T) {
	wrongCodes := []string{"", "___", "ZZ"}
	for _, code := range wrongCodes {
		c, err := Get(code)
		if err == nil {
			t.Errorf("Expected Error, got nil")
		}
		if len(c.Name()) > 0 {
			t.Errorf("Expected empty name, got %s", c.Name())
		}
		if len(c.Alpha2()) > 0 {
			t.Errorf("Expected empty alpha2 code, got %s", c.Alpha2())
		}
	}
}

func TestAR(t *testing.T) {
	c, err := Get("AR")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if c.Name() != "ARGENTINA" {
		t.Errorf("Expected name: ARGENTINA, got %s", c.Name())
	}
	if c.Alpha2() != "AR" {
		t.Errorf("Expected alpha2 code: AR, got %s", c.Alpha2())
	}
}

func TestMustGet(t *testing.T) {
	c := MustGet("AR")
	if c.Name() != "ARGENTINA" {
		t.Errorf("Expected name: ARGENTINA, got %s", c.Name())
	}
	if c.Alpha2() != "AR" {
		t.Errorf("Expected alpha2 code: AR, got %s", c.Alpha2())
	}

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	MustGet("ZZ")
}
