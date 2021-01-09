// Copyright (C) 2020  Germán Fuentes Capella

package credentials

import (
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type Credentials struct {
	slang.EKValue
}

const (
	credentialsType = "Credentials"
	credUsername    = "username"
	credPwd         = "pwd"
	credAccount     = "account"
)

func maskSensitive(field string, value starlark.Value) string {
	if field == credAccount || field == credPwd {
		return "\"*****\""
	}
	return value.String()
}

func New(username string, pwd string, account string) *Credentials {
	return &Credentials{
		slang.NewEKValue(
			credentialsType,
			[]string{credUsername, credPwd, credAccount},
			map[string]starlark.Value{
				credUsername: starlark.String(username),
				credPwd:      starlark.String(pwd),
				credAccount:  starlark.String(account),
			},
			map[string]slang.PreProcessFn{
				credUsername: slang.AssertString,
				credPwd:      slang.AssertString,
				credAccount:  slang.AssertString,
			},
			maskSensitive,
		),
	}
}
