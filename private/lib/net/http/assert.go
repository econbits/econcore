// Copyright (C) 2021  Germán Fuentes Capella

package http

import (
	"fmt"

	"go.starlark.net/starlark"
)

func AssertHTTPResponse(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(*HTTPResponse)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an HTTPResponse", v)
	}
	return v, nil
}

func AssertHeaderDict(v starlark.Value) (starlark.Value, error) {
	headerDict, ok := v.(*starlark.Dict)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an HTTPResponse", v)
	}
	for _, headerk := range headerDict.Keys() {
		_, ok := headerk.(starlark.String)
		if !ok {
			return nil, fmt.Errorf("HTTP header key '%v' is not a string", headerk)
		}
		vlist, found, err := headerDict.Get(headerk)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("HTTP header key '%v' is not in headers", headerk)
		}
		list, ok := vlist.(*starlark.List)
		if !ok {
			return nil, fmt.Errorf("HTTP header value '%v' is not a string list", vlist)
		}
		for i := 0; i < list.Len(); i++ {
			headerv := list.Index(i)
			_, ok := headerv.(starlark.String)
			if !ok {
				return nil, fmt.Errorf("HTTP header value '%v' is not of type string", vlist)
			}
		}
	}
	return headerDict, nil
}
