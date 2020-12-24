// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"go.starlark.net/starlark"
)

type Account struct {
	EKMValue
}

const (
	accountType = ekmValueType("Account")
	accName     = "name"
	accType     = "type"
	accIban     = "iban"
	accBic      = "bic"
)

// New function

func NewAccountValue(name starlark.String, typ starlark.String, iban starlark.String, bic starlark.String) *Account {
	return &Account{
		EKMValue{
			valueType: accountType,
			attrs:     []string{accName, accType, accIban, accBic},
			data: map[string]starlark.Value{
				accName: name,
				accType: typ,
				accIban: iban,
				accBic:  bic,
			},
			validatorsFn: map[string]validatorFunc{
				accName: isStringValue,
				accType: isStringValue,
				accIban: isStringValue,
				accBic:  isStringValue,
			},
			frozen: false,
			mask:   noMask,
		},
	}
}
