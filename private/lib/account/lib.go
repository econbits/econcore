// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/private/lib/account/ibanaccount"
	"github.com/econbits/econkit/private/lib/account/walletaccount"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "account",
		Fns: []*slang.Fn{
			ibanaccount.Fn,
			walletaccount.Fn,
		},
	}
)
