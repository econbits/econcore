// Copyright (C) 2021  Germán Fuentes Capella

package fin

import (
	"github.com/econbits/econkit/private/lib/fin/money"
	"github.com/econbits/econkit/private/lib/fin/transaction"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "fin",
		Fns: []*slang.Fn{
			money.Fn,
			transaction.Fn,
		},
	}
)
