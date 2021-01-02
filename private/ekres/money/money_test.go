//Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"math/big"
	"testing"

	"github.com/econbits/econkit/private/ekres/currency"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/money/"
	epilogue := starlark.StringDict{
		"money":    MoneyFn.Builtin(),
		"currency": currency.CurrencyFn.Builtin(),
	}
	testscript.TestingRun(t, dpath, epilogue, testscript.ExecScriptFn, testscript.Fail)
}

func TestZeroString(t *testing.T) {
	zero := big.NewInt(0)
	eur := currency.MustGet("EUR")

	m := New(zero, eur)
	if m.String() != "0.00EUR" {
		t.Errorf("Expected string: 0.00EUR; got: %s", m.String())
	}
}

func Test1EURCentString(t *testing.T) {
	one := big.NewInt(1)
	eur := currency.MustGet("EUR")

	m := New(one, eur)
	if m.String() != "0.01EUR" {
		t.Errorf("Expected string: 0.01EUR; got: %s", m.String())
	}
}

func TestEqual(t *testing.T) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	eur := currency.MustGet("EUR")

	m0 := New(zero, eur)
	m1 := New(one, eur)

	if !m0.Equal(m0) {
		t.Errorf("%v should be equal to itself", m0)
	}
	if m0.Equal(m1) {
		t.Errorf("%v should be not equal to %v", m0, m1)
	}
}
