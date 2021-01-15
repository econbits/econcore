// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/iban"
	"github.com/econbits/econkit/private/ekres/money"
	"github.com/econbits/econkit/private/ekres/session"
	"github.com/econbits/econkit/private/ekres/transaction"
	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

var (
	fns = []*slang.Fn{
		account.IbanFn,
		account.WalletFn,
		iban.IBANFn,
		money.MoneyFn,
		session.SessionFn,
		transaction.TransactionFn,
	}
)

func epilogue() starlark.StringDict {
	sd := starlark.StringDict{}
	for _, f := range fns {
		sd[f.Name] = f.Builtin()
	}
	for _, f := range universe.Lib.Fns {
		sd[f.Name] = f.Builtin()
	}
	return sd
}
