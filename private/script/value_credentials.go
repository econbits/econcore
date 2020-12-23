// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"go.starlark.net/starlark"
)

type Credentials struct {
	EKMValue
}

const (
	credentialsType = ekmValueType("Credentials")
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

// New function

func NewCredentials(username string, pwd string, account string) *Credentials {
	return &Credentials{
		EKMValue{
			valueType: credentialsType,
			attrs:     []string{credUsername, credPwd, credAccount},
			data: map[string]starlark.Value{
				credUsername: starlark.String(username),
				credPwd:      starlark.String(pwd),
				credAccount:  starlark.String(account),
			},
			validatorsFn: map[string]validatorFunc{
				credUsername: isStringValue,
				credPwd:      isStringValue,
				credAccount:  isStringValue,
			},
			frozen: false,
			mask:   maskSensitive,
		},
	}
}
