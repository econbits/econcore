// Copyright (C) 2021  Germán Fuentes Capella

package slang

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestEmptyLib(t *testing.T) {
	lib := Lib{Name: "TestLib", Fns: []*Fn{}}
	sd := lib.Load()
	if len(sd) != 0 {
		t.Fatalf("unexpected StringDict %v", sd)
	}
}

func TestNilLib(t *testing.T) {
	lib := Lib{Name: "TestLib", Fns: nil}
	sd := lib.Load()
	if len(sd) != 0 {
		t.Fatalf("unexpected StringDict %v", sd)
	}
}

func TestLib(t *testing.T) {
	fn := &Fn{
		Name: "TestFn",
		Callback: func(
			*starlark.Thread,
			*starlark.Builtin,
			starlark.Tuple,
			[]starlark.Tuple,
		) (starlark.Value, error) {
			return starlark.True, nil
		},
	}
	lib := Lib{Name: "TestLib", Fns: []*Fn{fn}}
	sd := lib.Load()
	if len(sd) != 1 {
		t.Fatalf("unexpected StringDict %v", sd)
	}
}
