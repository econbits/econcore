// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestSessionEasyMethods(t *testing.T) {
	s := Session{starlark.StringDict{}}

	s.Freeze()

	if s.Type() != "Session" {
		t.Errorf("Expected Session Type='Session'; got %s", s.Type())
	}

	if !strings.Contains(s.String(), "Session") {
		t.Errorf("Expected Session String to contain 'Session'; got %s", s.String())
	}
}

func TestSessionHashError(t *testing.T) {
	s := Session{starlark.StringDict{}}
	_, err := s.Hash()
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}
