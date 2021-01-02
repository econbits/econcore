// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"math/big"
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/currency"
	"go.starlark.net/starlark"
)

type Money struct {
	eklark.EKValue
}

const (
	moneyType = "Money"
	fAmount   = "amount"
	fCurrency = "currency"
	fnName    = "money"
)

var (
	errorClass = ekerrors.MustRegisterClass("MoneyError")
	MoneyFn    = &eklark.Fn{
		Name:     fnName,
		Callback: moneyFn,
	}
)

// amount gets cloned to avoid modifications outside of the Money interface
func New(amount *big.Int, curr *currency.Currency) *Money {
	zero := big.NewInt(0)
	amount = zero.Add(zero, amount)
	return &Money{
		eklark.NewEKValue(
			moneyType,
			[]string{
				fAmount,
				fCurrency,
			},
			map[string]starlark.Value{
				fAmount:   starlark.MakeBigInt(amount),
				fCurrency: curr,
			},
			map[string]eklark.ValidateFn{
				fAmount:   eklark.AssertInt,
				fCurrency: currency.AssertCurrency,
			},
			eklark.NoMaskFn,
		),
	}
}

func moneyFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var amount starlark.Int
	var currency *currency.Currency
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fAmount, &amount, fCurrency, &currency)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return New(amount.BigInt(), currency), nil
}

// It is safe to modify the returned *big.Int as it is a copy of the internal value
func (m *Money) Amount() *big.Int {
	intv := eklark.HasAttrsMustGetInt(m, fAmount)
	return intv.BigInt()
}

func (m *Money) Currency() *currency.Currency {
	return currency.HasAttrsMustGetCurrency(m, fCurrency)
}

func (m *Money) String() string {
	amountstr := eklark.HasAttrsMustGetInt(m, fAmount).String()
	currency := currency.HasAttrsMustGetCurrency(m, fCurrency)
	units := currency.Units()
	if units == 0 {
		return amountstr + currency.Code()
	}
	if amountstr == "0" {
		// units + 1 for: 0.00
		amountstr = strings.Repeat("0", units+1)
	}
	decPoint := len(amountstr) - units
	if decPoint < 0 {
		amountstr = strings.Repeat("0", -decPoint+1) + amountstr
		decPoint = len(amountstr) - units
	}
	lAmount := amountstr[0:decPoint]
	rAmount := amountstr[decPoint:]
	return lAmount + "." + rAmount + currency.Code()
}

func (m *Money) Equal(om *Money) bool {
	return m.Currency().Equal(om.Currency()) && m.Amount().Cmp(om.Amount()) == 0
}
