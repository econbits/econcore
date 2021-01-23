// Copyright (C) 2020  Germán Fuentes Capella

package assert

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	fnName     = "assert"
	defaultMsg = "Assertion Error"
)

var (
	Fn = &slang.Fn{
		Name:     fnName,
		Callback: assertCb,
	}
)

func assertCb(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var ok starlark.Bool
	var msg starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, "ok", &ok, "msg?", &msg)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{FormatError},
		)
	}
	if !ok.Truth() {
		if !msg.Truth() {
			return nil, ekerrors.New(
				errorClass,
				defaultMsg,
			)
		}
		return nil, ekerrors.New(
			errorClass,
			string(msg),
		)
	}
	return starlark.None, nil
}
