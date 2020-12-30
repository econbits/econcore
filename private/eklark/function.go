// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"

	"go.starlark.net/starlark"
)

type Callback func(*starlark.Thread, *starlark.Builtin, starlark.StringDict) (starlark.Value, error)

type Fn struct {
	Name     string
	ArgNames []string
	Callback Callback
	ArgError ErrorType
}

func (fn *Fn) validateNumArgs(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) error {
	if len(fn.ArgNames) < len(args)+len(kwargs) {
		return &EKError{
			FilePath:  ThreadMustGetFilePath(thread),
			Function:  builtin.Name(),
			ErrorType: fn.ArgError,
			Description: fmt.Sprintf(
				"%s() takes %d arguments but %d were provided",
				fn.Name,
				len(fn.ArgNames),
				len(args)+len(kwargs),
			),
		}
	}
	return nil
}

func (fn *Fn) value2string(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	v starlark.Value,
) (string, error) {
	str, ok := starlark.AsString(v)
	if !ok {
		return "", &EKError{
			FilePath:  ThreadMustGetFilePath(thread),
			Function:  builtin.Name(),
			ErrorType: fn.ArgError,
			Description: fmt.Sprintf(
				"%s(): invalid argument '%v'",
				fn.Name,
				v,
			),
		}
	}
	return str, nil
}

func (fn *Fn) popArgName(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	key string,
	argnames []string,
) ([]string, error) {
	for i, argname := range argnames {
		if key == argname {
			return append(argnames[:i], argnames[i+1:]...), nil
		}
	}
	return nil, &EKError{
		FilePath:  ThreadMustGetFilePath(thread),
		Function:  builtin.Name(),
		ErrorType: fn.ArgError,
		Description: fmt.Sprintf(
			"%s(): argument '%v' does not exist",
			fn.Name,
			key,
		),
	}
}

func (fn *Fn) starlarkCallback(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	err := fn.validateNumArgs(thread, builtin, args, kwargs)
	if err != nil {
		return nil, err
	}
	argnames := fn.ArgNames // avoid changing Fn ArgNames
	sdict := starlark.StringDict{}
	for _, value := range args {
		argname := argnames[0]
		argnames = argnames[1:]
		sdict[argname] = value
	}
	for _, kvpair := range kwargs {
		key, err := fn.value2string(thread, builtin, kvpair[0])
		if err != nil {
			return nil, err
		}
		argnames, err = fn.popArgName(thread, builtin, key, argnames)
		if err != nil {
			return nil, err
		}
		sdict[key] = kvpair[1]
	}
	for _, argname := range argnames {
		sdict[argname] = starlark.None
	}
	return fn.Callback(thread, builtin, sdict)
}

func (fn *Fn) Builtin() *starlark.Builtin {
	return starlark.NewBuiltin(
		fn.Name,
		fn.starlarkCallback,
	)
}
