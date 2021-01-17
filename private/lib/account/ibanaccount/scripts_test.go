// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"testing"

	"github.com/econbits/econkit/private/lib/iso/bic"
	"github.com/econbits/econkit/private/lib/iso/iban"
	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "000_smalltests/lib/account/iban_account/"
	testscript.TestingRun(
		t,
		dpath,
		universe.Lib.Load(),
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			sd := starlark.StringDict{
				bic.Fn.Name:  bic.Fn.Builtin(),
				Fn.Name:      Fn.Builtin(),
				iban.Fn.Name: iban.Fn.Builtin(),
			}
			return sd, nil
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
