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

func assertCb(thread *starlark.Thread, builtin *starlark.Builtin, sdict eklark.StringDict) (starlark.Value, error) {
	ok, err := sdict.GetBool(argOk)
	if err != nil {
		return nil, &eklark.EKError{
			FilePath:    eklark.MustGetFilePath(thread),
			Function:    builtin.Name(),
			ErrorType:   fnErrorType,
			Description: err.Error(),
		}
	}
	if !ok {
		msg, err := sdict.GetStringOr(argMsg, defaultMsg)
		if err != nil {
			return nil, &eklark.EKError{
				FilePath:    eklark.MustGetFilePath(thread),
				Function:    builtin.Name(),
				ErrorType:   fnErrorType,
				Description: err.Error(),
			}
		}
		return nil, &eklark.EKError{
			FilePath:    eklark.MustGetFilePath(thread),
			Function:    builtin.Name(),
			ErrorType:   fnErrorType,
			Description: msg,
		}
	}
	return starlark.None, nil
}
