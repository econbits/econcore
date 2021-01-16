// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"fmt"
	"math/big"

	"github.com/econbits/econkit/private/ekerrors"
)

var (
	mathError = ekerrors.MustRegisterClass("MoneyMathError")
)

func (m *Money) Add(om *Money) (*Money, error) {
	currency := m.Currency()
	if !currency.Equal(om.Currency()) {
		return nil, ekerrors.New(
			mathError,
			fmt.Sprintf("%v and %v can't be added: different currency", m, om),
		)
	}
	amount := big.NewInt(0).Add(m.Amount(), om.Amount())
	return New(amount, currency), nil
}

func (m *Money) Sub(om *Money) (*Money, error) {
	currency := m.Currency()
	if !currency.Equal(om.Currency()) {
		return nil, ekerrors.New(
			mathError,
			fmt.Sprintf("%v and %v can't be subtracted: different currency", m, om),
		)
	}
	amount := m.Amount()
	amount = amount.Sub(amount, om.Amount())
	return New(amount, currency), nil
}
