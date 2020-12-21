//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

func builtin_session(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	if len(args) > 0 {
		return Session{}, newBuiltinError(
			b.Name(),
			SessionError,
			fmt.Sprintf("unnamed arguments are not allowed: %v", args),
		)
	}
	s := Session{d: starlark.StringDict{}}
	if len(kwargs) > 0 {
		for _, tuple := range kwargs {
			s.SetKey(tuple[0], tuple[1])
		}
	}
	return s, nil
}

func builtin_account(
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
					return Account{}, newBuiltinError(
						b.Name(),
						AccountError,
						fmt.Sprintf("'%s' is a non-string param", args[i]),
					)
				}
				params[argname] = argvalue
			}
		}
	}
	for _, tuple := range kwargs {
		argname, ok := starlark.AsString(tuple[0])
		if !ok {
			return Account{}, newBuiltinError(
				b.Name(),
				AccountError,
				fmt.Sprintf("'%s' is a non-string keyword argument", tuple[0]),
			)
		}
		if params[argname].Len() > 0 {
			return Account{}, newBuiltinError(
				b.Name(),
				AccountError,
				fmt.Sprintf("'%s' appears as positional argument and keyword argument", argname),
			)
		}
		argvalue, ok := tuple[1].(starlark.String)
		if !ok {
			return Account{}, newBuiltinError(
				b.Name(),
				AccountError,
				fmt.Sprintf("'%s' is a non-string param", tuple[1]),
			)
		}
		params[argname] = argvalue
	}
	return Account{
		name: params["name"],
		typ:  params["type"],
		iban: params["iban"],
		bic:  params["bic"],
	}, nil
}
