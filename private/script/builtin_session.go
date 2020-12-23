//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

func builtinSession(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	if len(args) > 0 {
		return nil, newBuiltinError(
			b.Name(),
			SessionError,
			fmt.Sprintf("unnamed arguments are not allowed: %v", args),
		)
	}
	s := NewSession()
	if len(kwargs) > 0 {
		for _, tuple := range kwargs {
			s.SetKey(tuple[0], tuple[1])
		}
	}
	return s, nil
}
