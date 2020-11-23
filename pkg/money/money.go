//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"github.com/econbits/econcore/pkg/currency"
)

type Money struct {
	amount   uint64
	currency currency.Currency
}

// Returns the currency for Money
func (m Money) Currency() currency.Currency {
	return m.currency
}
