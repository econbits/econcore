// Copyright (C) 2021  Germán Fuentes Capella

package account

const (
	KindChecking   = "checking"
	KindSavings    = "savings"
	KindLoan       = "loan"
	KindWallet     = "wallet"
	KindCreditCard = "credit card"
)

var (
	account_types = []string{
		KindChecking,
		KindSavings,
		KindLoan,
		KindWallet,
		KindCreditCard,
	}
)
