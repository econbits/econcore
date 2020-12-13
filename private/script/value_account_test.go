//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"strings"
	"testing"
)

func TestAccountEasyMethods(t *testing.T) {
	a := Account{}

	a.Freeze()

	if a.Type() != "Account" {
		t.Errorf("Expected Account Type='Account'; got %s", a.Type())
	}

	if !strings.Contains(a.String(), "Account") {
		t.Errorf("Expected Account String to contain 'Account'; got %s", a.String())
	}

	if a.Truth() {
		t.Errorf("Expected Account Truth=false; got true")
	}
}

func TestAccountHashError(t *testing.T) {
	a := Account{}
	_, err := a.Hash()
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}
