// Copyright (C) 2020  Germán Fuentes Capella

package slang

import (
	"testing"
)

func TestScriptId(t *testing.T) {
	id := ScriptId("id.ekm")
	if id != "id" {
		t.Fatalf("expected 'id'; got '%s'", id)
	}
}
