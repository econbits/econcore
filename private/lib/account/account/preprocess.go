// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

func MatchKind(v starlark.Value, subset []string) (starlark.Value, error) {
	v, err := slang.AssertString(v)
	if err != nil {
		return nil, err
	}

	reqType, _ := starlark.AsString(v)
	for _, subsetType := range subset {
		if subsetType == reqType {
			// if there is a match, further validate that the types slice
			// contains valid account types
			for _, defType := range account_types {
				if defType == reqType {
					return v, nil
				}
			}
		}
	}

	return nil, ekerrors.New(
		errorClass,
		fmt.Sprintf("'%s' is not a valid account type", reqType),
	)
}

func matchAccountKind(v starlark.Value) (starlark.Value, error) {
	return MatchKind(v, account_types)
}
