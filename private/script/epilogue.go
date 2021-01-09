// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/assert"
	"github.com/econbits/econkit/private/ekres/bic"
	"github.com/econbits/econkit/private/ekres/country"
	"github.com/econbits/econkit/private/ekres/currency"
	"github.com/econbits/econkit/private/ekres/datetime"
	"github.com/econbits/econkit/private/ekres/iban"
	"github.com/econbits/econkit/private/ekres/money"
	"github.com/econbits/econkit/private/ekres/session"
	"github.com/econbits/econkit/private/ekres/transaction"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

var (
	fns = []*slang.Fn{
		account.IbanFn,
		account.WalletFn,
		assert.AssertFn,
		bic.BICFn,
		country.CountryFn,
		currency.CurrencyFn,
		datetime.DateTimeFn,
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
	return sd
}
