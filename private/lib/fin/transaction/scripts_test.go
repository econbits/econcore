// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"fmt"
	"testing"

	"github.com/econbits/econkit/private/lib/account"
	"github.com/econbits/econkit/private/lib/datetime"
	"github.com/econbits/econkit/private/lib/fin/money"
	"github.com/econbits/econkit/private/lib/iso"
	"github.com/econbits/econkit/private/slang"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../../test/ekm/vdefault/000_smalltests/ekres/transaction/"
	testscript.TestingRun(
		t,
		dpath,
		starlark.StringDict{},
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			libs := []*slang.Lib{
				account.Lib,
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
