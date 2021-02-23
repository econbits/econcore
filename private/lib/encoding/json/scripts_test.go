// Copyright (C) 2021  Germán Fuentes Capella

package json

import (
	"fmt"
	"testing"

	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "000_smalltests/lib/encoding/json/"
	testscript.TestingRun(
		t,
		dpath,
		universe.Lib.Load(),
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			if module == "encoding" {
				sd := starlark.StringDict{
					Fn.Name: Fn.Builtin(),
				}
				return sd, nil
			}
			return nil, fmt.Errorf("unknown module: %s", module)
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
