// Copyright (C) 2021  Germán Fuentes Capella

package slang

import (
	"go.starlark.net/starlark"
)

type Lib struct {
	Name string
	Fns  []*Fn
}

func (lib *Lib) Load() starlark.StringDict {
	sd := starlark.StringDict{}
	if lib.Fns != nil {
		for _, fn := range lib.Fns {
			sd[fn.Name] = fn.Builtin()
		}
	}
	return sd
}
