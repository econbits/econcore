// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"

	"github.com/econbits/econkit/private/ekres/iban"
	"github.com/econbits/econkit/private/lib/iso/bic"
	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/account/"
	epilogue := starlark.StringDict{
		IbanFn.Name:   IbanFn.Builtin(),
		WalletFn.Name: WalletFn.Builtin(),
		//
		iban.IBANFn.Name: iban.IBANFn.Builtin(),
	}
	for name, builtin := range universe.Lib.Load() {
		epilogue[name] = builtin
	}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			sd := starlark.StringDict{
				bic.BICFn.Name: bic.BICFn.Builtin(),
			}
			return sd, nil
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
