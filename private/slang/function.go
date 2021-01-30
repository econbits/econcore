// Copyright (C) 2020  Germán Fuentes Capella

package slang

import (
	"errors"
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"go.starlark.net/starlark"
)

var (
	fnErrorClass = ekerrors.MustRegisterClass("ParsingFunctionArgsError")
)

type Callback func(*starlark.Thread, *starlark.Builtin, starlark.Tuple, []starlark.Tuple) (starlark.Value, error)

type Fn struct {
	Name     string
	Callback Callback
}

func (fn *Fn) Builtin() *starlark.Builtin {
	return starlark.NewBuiltin(
		fn.Name,
		validatedCb(fn.Callback),
	)
}

func validatedCb(cb Callback) Callback {
	return func(
		thread *starlark.Thread,
		builtin *starlark.Builtin,
		args starlark.Tuple,
		kwargs []starlark.Tuple,
	) (starlark.Value, error) {
		for _, pair := range kwargs {
			_, ok := starlark.AsString(pair[0])
			if !ok {
				return nil, ekerrors.New(
					fnErrorClass,
					fmt.Sprintf("attribute '%v' must be of type string, not '%T'", pair[0], pair[0]),
				)
			}
		}
		val, err := cb(thread, builtin, args, kwargs)
		if err != nil {
			var ekerr *ekerrors.EKError
			if errors.As(err, &ekerr) {
				ekerr.LinkCS(thread.CallStack())
			}
		}
		return val, err
	}
}
