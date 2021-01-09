// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

const (
	hDescription = "DESCRIPTION"
	hURL         = "URL"
	hAuthors     = "AUTHORS"
	hLicense     = "LICENSE"
)

var (
	stringHeaders = []string{
		hDescription,
		hURL,
		hLicense,
	}
	listHeaders = []string{
		hAuthors,
	}
)

const (
	defDescription = "Module under construction"
	defUrl         = "https://econbits.org/"
	defLicense     = ""
)

var (
	defAuthors = []string{}
)

func validateGlobalVars(g starlark.StringDict) error {
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
