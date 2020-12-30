// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"go.starlark.net/starlark"
)

func Exec(thread *starlark.Thread, filePath string, epilogue starlark.StringDict) (starlark.StringDict, error) {
	globals, err := starlark.ExecFile(thread, filePath, nil, epilogue)
	if err != nil {
		evalerr, ok := err.(*starlark.EvalError)
		if ok {
			ekerr, ok := evalerr.Unwrap().(*EKError)
			if ok {
				return nil, ekerr
			}
		}
		return nil, err
	}
	return globals, nil
}
