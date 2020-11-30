//Copyright (C) 2020  Germán Fuentes Capella

package country

import (
	"testing"
)

func TestIsEqual(t *testing.T) {
	c, err := Get("AR")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	if !c.IsEqual(c) {
		t.Errorf("'%v' is not equal to itself", c)
	}
}

func TestIsNotEqual(t *testing.T) {
	c1, err := Get("AR")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	c2, err := Get("DE")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	if c1.IsEqual(c2) {
		t.Errorf("'%v' is equal to '%v'", c1, c2)
	}
}
