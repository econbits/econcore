// Copyright (C) 2021  Germán Fuentes Capella

package ekm

import (
	"testing"

	"github.com/econbits/econkit/private/lib/datetime"
	"go.starlark.net/starlark"
)

func TestLoadMissingImport(t *testing.T) {
	thread := &starlark.Thread{Name: "TestThread", Load: load}
	_, err := load(thread, "import-does-not-exist")
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestLoadDateTime(t *testing.T) {
	thread := &starlark.Thread{Name: "TestThread", Load: load}
	sd, err := load(thread, datetime.Lib.Name)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	for _, fn := range datetime.Lib.Fns {
		_, ok := sd[fn.Name]
		if !ok {
			t.Fatalf("%s is not loaded in %v", fn.Name, datetime.Lib)
		}
	}
}
