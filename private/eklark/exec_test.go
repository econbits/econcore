// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestExecOKScript(t *testing.T) {
	filePath := "../../test/ekm/vdefault/000_smalltests/eklark/empty.ekm"
	thread := NewThread("Test")
	_, err := Exec(thread, filePath, starlark.StringDict{})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestExecFailScript(t *testing.T) {
	filePath := "../../test/ekm/vdefault/000_smalltests/eklark/fail.ekm"
	thread := NewThread("Test")
	_, err := Exec(thread, filePath, starlark.StringDict{})
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
