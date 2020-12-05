//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"github.com/econbits/econkit/pkg/currency"
)

func New(amount uint64, currencyCode string) (Money, error) {
	c, err := currency.Get(currencyCode)
	if err != nil {
		return Money{}, CurrencyNotFoundError(err.Error())
	}
	return Money{amount: amount, currency: c}, nil
}

func Zero(currencyCode string) (Money, error) {
	return New(0, currencyCode)
}

func MustNew(amount uint64, currencyCode string) Money {
	m, err := New(amount, currencyCode)
	if err != nil {
		panic(err.Error())
	}
	return m
}

func MustZero(currencyCode string) Money {
	return MustNew(0, currencyCode)
}
