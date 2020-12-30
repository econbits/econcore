// Copyright (C) 2020  Germán Fuentes Capella

package assert

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

const (
	fnName     = "assert"
	argOk      = "ok"
	argMsg     = "msg"
	defaultMsg = "Assertion Error"
)

var (
	fnArgNames   = []string{argOk, argMsg}
	fnErrorClass = ekerrors.MustRegisterClass("AssertionError")
	AssertFn     = &eklark.Fn{
		Name:          fnName,
		ArgNames:      fnArgNames,
		Callback:      assertCb,
		ArgErrorClass: fnErrorClass,
	}
)

func assertCb(thread *starlark.Thread, builtin *starlark.Builtin, sdict starlark.StringDict) (starlark.Value, error) {
	ok, err := eklark.StringDictGetBool(sdict, argOk)
	if err != nil {
		return nil, ekerrors.Wrap(
			fnErrorClass,
			err.Error(),
			err,
		)
	}
	if !ok {
		msg, err := eklark.StringDictGetStringOr(sdict, argMsg, defaultMsg)
		if err != nil {
			return nil, ekerrors.Wrap(
				fnErrorClass,
				err.Error(),
				err,
			)
		}
		return nil, ekerrors.New(
			fnErrorClass,
			msg,
		)
	}
	return starlark.None, nil
}
