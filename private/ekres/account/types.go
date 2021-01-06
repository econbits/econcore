// Copyright (C) 2021  Germán Fuentes Capella

package account

const (
	checkingType   = "checking"
	savingsType    = "savings"
	loanType       = "loan"
	walletType     = "wallet"
	creditCardType = "credit card"
)

var (
	account_types = []string{
		checkingType,
		savingsType,
		loanType,
		walletType,
		creditCardType,
	}
	iban_account_types = []string{
		checkingType,
		savingsType,
	}
	wallet_account_types = []string{
		walletType,
	}
)
