// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

func New(fpath string) (Script, error) {
	name := scriptid(fpath)
	thread := &starlark.Thread{Name: name}
	globals, err := starlark.ExecFile(thread, fpath, nil, epilogue)
	if err != nil {
		return Script{}, ScriptError{
			fpath:     fpath,
			function:  "load",
			errorType: LoadError,
			text:      err.Error(),
		}
	}
	err = validateGlobals(globals)
	if err != nil {
		return Script{}, ScriptError{
			fpath:     fpath,
			function:  "globals",
			errorType: ReservedVarError,
			text:      err.Error(),
		}
	}
	return Script{tn: thread, fpath: fpath, globals: globals}, nil
}

func validateGlobals(g starlark.StringDict) error {
	for _, h := range stringHeaders {
		field, ok := g[h]
		if ok {
			_, ok := field.(starlark.String)
			if !ok {
				return fmt.Errorf("'%s' must be of type string", h)
			}
		}
	}
	for _, h := range listHeaders {
		field, ok := g[h]
		if ok {
			lfield, ok := field.(*starlark.List)
			if !ok {
				return fmt.Errorf("'%s' must be of type list", h)
			}
			for i := 0; i < lfield.Len(); i++ {
				sfield := lfield.Index(i)
				_, ok := sfield.(starlark.String)
				if !ok {
					return fmt.Errorf("items in '%s' must be of type string", h)
				}
			}
		}
	}
	return nil
}
