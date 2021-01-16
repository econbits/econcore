// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"fmt"
	"testing"

	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/lib/datetime"
	"github.com/econbits/econkit/private/lib/fin/money"
	"github.com/econbits/econkit/private/lib/iso"
	"github.com/econbits/econkit/private/slang"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../../test/ekm/vdefault/000_smalltests/ekres/transaction/"
	fns := []*slang.Fn{
		account.WalletFn,
	}
	epilogue := starlark.StringDict{}
	for _, fn := range fns {
		epilogue[fn.Name] = fn.Builtin()
	}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			libs := []*slang.Lib{
				datetime.Lib,
				&slang.Lib{
					Name: "fin",
					Fns: []*slang.Fn{
						money.Fn,
						Fn,
					},
				},
				iso.Lib,
			}
			for _, lib := range libs {
				if module == lib.Name {
					return lib.Load(), nil
				}
			}
			return nil, fmt.Errorf("unknown module: %s", module)
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
