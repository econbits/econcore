// Copyright (C) 2020  Germán Fuentes Capella

package assert

import (
	"testing"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../../test/ekm/vdefault/000_smalltests/ekres/assert/"
	epilogue := starlark.StringDict{Fn.Name: Fn.Builtin()}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
