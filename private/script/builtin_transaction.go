//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

func builtinTransactions(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	if len(args) > 0 {
		return nil, ScriptError{
			scriptName: "",
			function:   b.Name(),
			errorType:  TransactionError,
			text:       fmt.Sprintf("unnamed arguments are not allowed: %v", args),
		}
	}
	if len(kwargs) == 0 {
		return nil, ScriptError{
			scriptName: "",
			function:   b.Name(),
			errorType:  TransactionError,
			text:       "Missing mandatory named arguments: {src_iban, dst_iban, amount, currency, value_date}",
		}
	}
	t := NewTransaction()
	if len(kwargs) > 0 {
		for _, tuple := range kwargs {
			key, ok := starlark.AsString(tuple[0])
			if !ok {
				return nil, ScriptError{
					scriptName: "",
					function:   b.Name(),
					errorType:  TransactionError,
					text:       fmt.Sprintf("kwargs indexed with type: %T; value: %v", tuple[0], tuple[0]),
				}
			}
			err := t.SetField(key, tuple[1])
			if err != nil {
				return nil, ScriptError{
					scriptName: "",
					function:   b.Name(),
					errorType:  TransactionError,
					text:       err.Error(),
				}
			}
		}
	}
	return t, nil
}
