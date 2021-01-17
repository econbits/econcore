// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"github.com/econbits/econkit/private/lib/account/account"
)

var (
	iban_account_types = []string{
		account.KindChecking,
		account.KindSavings,
	}
)
