//Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"testing"

	"github.com/econbits/econkit/private/lib/iso/currency"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../../test/ekm/vdefault/000_smalltests/ekres/money/"
	testscript.TestingRun(
		t,
		dpath,
		starlark.StringDict{},
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			sd := starlark.StringDict{
				MoneyFn.Name:             MoneyFn.Builtin(),
				currency.CurrencyFn.Name: currency.CurrencyFn.Builtin(),
			}
			return sd, nil
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
