// Copyright (C) 2021  Germán Fuentes Capella

package walletaccount

import (
	"testing"

	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "000_smalltests/lib/account/wallet_account/"
	testscript.TestingRun(
		t,
		dpath,
		universe.Lib.Load(),
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			sd := starlark.StringDict{
				Fn.Name: Fn.Builtin(),
			}
			return sd, nil
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
