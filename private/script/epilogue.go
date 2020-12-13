//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"go.starlark.net/starlark"
)

var (
	epilogue = starlark.StringDict{
		"session": starlark.NewBuiltin("session", builtin_session),
		"account": starlark.NewBuiltin("account", builtin_account),
	}
)
