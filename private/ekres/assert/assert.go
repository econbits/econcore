// Copyright (C) 2020  Germán Fuentes Capella

package assert

import (
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

const (
	fnName      = "assert"
	fnErrorType = eklark.ErrorType("AssertionError")
	argOk       = "ok"
	argMsg      = "msg"
	defaultMsg  = "Assertion Error"
)

var (
	fnArgNames = []string{argOk, argMsg}
	AssertFn   = &eklark.Fn{
		Name:     fnName,
		ArgNames: fnArgNames,
		Callback: assertCb,
		ArgError: fnErrorType,
	}
)

func assertCb(thread *starlark.Thread, builtin *starlark.Builtin, sdict starlark.StringDict) (starlark.Value, error) {
	ok, err := eklark.StringDictGetBool(sdict, argOk)
	if err != nil {
		return nil, &eklark.EKError{
			FilePath:    eklark.ThreadMustGetFilePath(thread),
			Function:    builtin.Name(),
			ErrorType:   fnErrorType,
			Description: err.Error(),
		}
	}
	if !ok {
		msg, err := eklark.StringDictGetStringOr(sdict, argMsg, defaultMsg)
		if err != nil {
			return nil, &eklark.EKError{
				FilePath:    eklark.ThreadMustGetFilePath(thread),
				Function:    builtin.Name(),
				ErrorType:   fnErrorType,
				Description: err.Error(),
			}
		}
		return nil, &eklark.EKError{
			FilePath:    eklark.ThreadMustGetFilePath(thread),
			Function:    builtin.Name(),
			ErrorType:   fnErrorType,
			Description: msg,
		}
	}
	return starlark.None, nil
}
