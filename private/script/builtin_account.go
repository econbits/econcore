// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

func builtinAccounts(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	params := map[string]starlark.String{
		"name": starlark.String(""),
		"type": starlark.String(""),
		"iban": starlark.String(""),
		"bic":  starlark.String(""),
	}
	argdef := []string{"name", "type", "iban", "bic"}
	if len(args) > 0 {
		for i, argname := range argdef {
			if len(args) >= (i + 1) {
				argvalue, ok := args[i].(starlark.String)
				if !ok {
					return nil, ScriptError{
						fpath:     "",
						function:  b.Name(),
						errorType: AccountError,
						text:      fmt.Sprintf("'%s' is a non-string param", args[i]),
					}
				}
				params[argname] = argvalue
			}
		}
	}
	for _, tuple := range kwargs {
		argname, ok := starlark.AsString(tuple[0])
		if !ok {
			return nil, ScriptError{
				fpath:     "",
				function:  b.Name(),
				errorType: AccountError,
				text:      fmt.Sprintf("'%s' is a non-string keyword argument", tuple[0]),
			}
		}
		if params[argname].Len() > 0 {
			return nil, ScriptError{
				fpath:     "",
				function:  b.Name(),
				errorType: AccountError,
				text:      fmt.Sprintf("'%s' appears as positional argument and keyword argument", argname),
			}
		}
		argvalue, ok := tuple[1].(starlark.String)
		if !ok {
			return nil, ScriptError{
				fpath:     "",
				function:  b.Name(),
				errorType: AccountError,
				text:      fmt.Sprintf("'%s' is a non-string param", tuple[1]),
			}
		}
		params[argname] = argvalue
	}
	return NewAccountValue(params["name"], params["type"], params["iban"], params["bic"]), nil
}
