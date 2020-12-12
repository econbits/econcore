//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"
	"path/filepath"

	"go.starlark.net/starlark"
)

func fname(fpath string) string {
	name := filepath.Base(fpath)
	ext := filepath.Ext(name)
	return name[0 : len(name)-len(ext)]
}

func validateGlobals(name string, g starlark.StringDict) error {
	for _, h := range stringHeaders {
		field, ok := g[h]
		if ok {
			_, ok := field.(starlark.String)
			if !ok {
				return fmt.Errorf("[%s][%s] %s must be of type string", name, h, h)
			}
		}
	}
	for _, h := range listHeaders {
		field, ok := g[h]
		if ok {
			lfield, ok := field.(*starlark.List)
			if !ok {
				return fmt.Errorf("[%s][%s] %s must be of type list", name, h, h)
			}
			for i := 0; i < lfield.Len(); i++ {
				sfield := lfield.Index(i)
				_, ok := sfield.(starlark.String)
				if !ok {
					return fmt.Errorf("[%s][%s] items in %s must be of type string", name, h, h)
				}
			}
		}
	}
	return nil
}
