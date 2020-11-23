//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"github.com/econbits/econcore/pkg/currency"
)

func New(amount uint64, currencyCode string) (Money, error) {
	c, err := currency.Get(currencyCode)
	if err != nil {
		return noMoney, CurrencyNotFoundError(err.Error())
	}
	return Money{amount: amount, currency: c}, nil
}
