// Copyright (C) 2021  Germán Fuentes Capella

package base64

import (
	pbase64 "encoding/base64"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	encFnNameme = "base64_encode"
)

var (
	EncFn = &slang.Fn{
		Name:     encFnNameme,
		Callback: encodeFn,
	}
)

func encodeFn(
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
	str := pbase64.StdEncoding.EncodeToString([]byte(text))
	return starlark.String(str), nil
}
