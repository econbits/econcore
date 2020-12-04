//Copyright (C) 2020  Germán Fuentes Capella

package account

type Type uint16

const (
	TypeCreditCard       = Type(1)
	TypeFixedTermDeposit = Type(2)
	TypeGiro             = Type(3)
	TypeLoan             = Type(4)
	TypePortfolio        = Type(5)
	TypeSavings          = Type(6)
	TypeCrypto           = Type(7)
)
