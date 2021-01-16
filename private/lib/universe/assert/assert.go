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
	fnErrorClass = ekerrors.MustRegisterClass("AssertionError")
	Fn           = &slang.Fn{
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
			fnErrorClass,
			err.Error(),
			err,
		)
	}
	if !ok.Truth() {
		if !msg.Truth() {
			return nil, ekerrors.New(
				fnErrorClass,
				defaultMsg,
			)
		}
		return nil, ekerrors.New(
			fnErrorClass,
			string(msg),
		)
	}
	return starlark.None, nil
}
