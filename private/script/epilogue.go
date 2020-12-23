//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"go.starlark.net/starlark"
)

var (
	epilogue = starlark.StringDict{
		"session":     starlark.NewBuiltin("session", builtinSession),
		"account":     starlark.NewBuiltin("account", builtinAccounts),
		"transaction": starlark.NewBuiltin("transaction", builtinTransactions),
	}
)
