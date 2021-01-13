// Copyright (C) 2021  Germán Fuentes Capella

package country

import (
	"testing"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/country/"
	fn := CountryFn
	epilogue := starlark.StringDict{fn.Name: fn.Builtin()}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}

func TestDE(t *testing.T) {
	code := "DE"
	c := MustGet(code)

	if c.Alpha2() != code {
		t.Fatalf("expected %s; got %s", code, c.Alpha2())
	}
	if c.String() != code {
		t.Fatalf("expected %s; got %s", code, c.String())
	}
	name := "GERMANY"
	if c.Name() != name {
		t.Fatalf("expected %s; got %s", name, c.Name())
	}
}

func TestMustGetPanic(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	MustGet("blabla")
}

func TestGetInvalidAlpha2(t *testing.T) {
	_, err := Get("blablabla")
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
