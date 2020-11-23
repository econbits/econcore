//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"testing"
)

func TestIsEqual(t *testing.T) {
	m1, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	m2 := MustParse("1.00EUR")
	if !m1.IsEqual(m2) {
		t.Errorf("'%v' is not equal to '%v'", m1, m2)
	}
}

func TestIsNotEqual(t *testing.T) {
	m1, err := New(100, "EUR")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	m2 := MustParse("2.00EUR")
	if m1.IsEqual(m2) {
		t.Errorf("'%v' is equal to '%v'", m1, m2)
	}
}
