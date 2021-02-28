// Copyright (C) 2021  Germán Fuentes Capella

package base64

import (
	pbase64 "encoding/base64"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	decFnNameme = "base64_decode"
)

var (
	DecFn = &slang.Fn{
		Name:     decFnNameme,
		Callback: decodeFn,
	}
)

func decodeFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var text starlark.String
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		"text", &text,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	data, err := pbase64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	return starlark.Bytes(data), nil
}
