// Copyright (C) 2020  Germán Fuentes Capella

package currency

import (
	"testing"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/currency/"
	epilogue := starlark.StringDict{"currency": CurrencyFn.Builtin()}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}

func TestEUR(t *testing.T) {
	curr, err := Get("EUR")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	name, code, id, units := "Euro", "EUR", 978, 2
	if curr.Name() != name {
		t.Fatalf("expected %s; found %s", name, curr.Name())
	}
	if curr.Code() != code {
		t.Fatalf("expected %s; found %s", code, curr.Code())
	}
	if curr.Id() != id {
		t.Fatalf("expected %d; found %d", id, curr.Id())
	}
	if curr.Units() != units {
		t.Fatalf("expected %d; found %d", units, curr.Units())
	}
}

func TestEqual(t *testing.T) {
	eur, err := Get("EUR")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	usd, err := Get("USD")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if eur.Equal(usd) {
		t.Fatalf("ERROR %s == %s", eur.Code(), usd.Code())
	}
	if !eur.Equal(eur) {
		t.Fatalf("ERROR %s != %s", eur.Code(), eur.Code())
	}
}

func TestMustGet(t *testing.T) {
	MustGet("EUR")
}

func TestMustGetPanic(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	MustGet("blabla")
}
