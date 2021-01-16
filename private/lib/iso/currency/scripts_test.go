// Copyright (C) 2021  Germán Fuentes Capella

package currency

import (
	"testing"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../../test/ekm/vdefault/000_smalltests/ekres/currency/"
	testscript.TestingRun(
		t,
		dpath,
		starlark.StringDict{},
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
