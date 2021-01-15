// Copyright (C) 2021  Germán Fuentes Capella

package ekm

import (
	"testing"

	"github.com/econbits/econkit/private/lib/datetime"
	"github.com/econbits/econkit/private/lib/iso"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

func TestLoadMissingImport(t *testing.T) {
	thread := &starlark.Thread{Name: "TestThread", Load: load}
	_, err := load(thread, "import-does-not-exist")
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestLoadLib(t *testing.T) {
	libs := map[string][]*slang.Fn{
		datetime.Lib.Name: datetime.Lib.Fns,
		iso.Lib.Name:      iso.Lib.Fns,
	}

	for libname, fns := range libs {
		thread := &starlark.Thread{Name: "TestThread", Load: load}
		sd, err := load(thread, libname)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		for _, fn := range fns {
			_, ok := sd[fn.Name]
			if !ok {
				t.Fatalf("%s is not loaded in %v", fn.Name, libname)
			}
		}
	}
}
