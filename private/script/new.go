//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

func New(fpath string) (Script, error) {
	name := fname(fpath)
	thread := &starlark.Thread{Name: name}
	globals, err := starlark.ExecFile(thread, fpath, nil, epilogue)
	if err != nil {
		return Script{}, fmt.Errorf("[%s][load] %v", name, err)
	}
	err = validateGlobals(name, globals)
	if err != nil {
		return Script{}, err
	}
	return Script{tn: thread, fpath: fpath, globals: globals}, nil
}
