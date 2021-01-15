// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"fmt"
	"math/big"
	"regexp"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/lib/iso/currency"
)

var (
	moneyParseAmountError   = ekerrors.MustRegisterClass("MoneyParseAmountError")
	moneyParseCurrencyError = ekerrors.MustRegisterClass("MoneyParseCurrencyError")
	moneyRe                 = regexp.MustCompile(`(?i)^\s*(?P<lh>[0-9]+)(\.(?P<rh>[0-9]+)){0,1}\s*(?P<curr>[a-z]{3}){0,1}\s*$`)
)

func splitComponents(moneystr string) (string, int, string) {
	groupNames := moneyRe.SubexpNames()
	matches := moneyRe.FindStringSubmatch(moneystr)

	amountstr := ""
	units := 0
	currencyCode := ""
	for i, value := range matches {
		if i == 0 {
			// this is the whole str
			continue
		}
		if groupNames[i] == "lh" {
			amountstr = value + amountstr
		} else if groupNames[i] == "rh" {
			amountstr = amountstr + value
			units = len(value)
		} else if groupNames[i] == "curr" {
			currencyCode = value
		}
	}
	return amountstr, units, currencyCode
}

func validateComponents(moneystr string, amountstr string, currencyCode string) error {
	if len(amountstr) == 0 {
		return ekerrors.New(
			moneyParseAmountError,
			fmt.Sprintf("Amount not found in '%s'", moneystr),
		)
	}
	if len(currencyCode) == 0 {
		return ekerrors.New(
			moneyParseCurrencyError,
			fmt.Sprintf("Currency not found in '%s'", moneystr),
		)
	}
	return nil
}

func shiftToCurrencyUnits(amount *big.Int, currency *currency.Currency, units int) error {
	if currency.Units() < units {
		return ekerrors.New(
			moneyParseCurrencyError,
			fmt.Sprintf("'%d' are too many decimal points for [%s:%d]", units, currency.Code(), currency.Units()),
		)
	}

	shift := int64(1)
	currencyUnits := currency.Units()
	for ; units < currencyUnits; units++ {
		shift = shift * int64(10)
	}
	amount.Mul(amount, big.NewInt(shift))
	return nil
}

func Parse(moneystr string) (*Money, error) {
	amountstr, units, currencyCode := splitComponents(moneystr)
	err := validateComponents(moneystr, amountstr, currencyCode)
	if err != nil {
		return nil, err
	}

	amount := big.NewInt(0)
	amount, ok := amount.SetString(amountstr, 10)
	if !ok {
		// I don't think this erros is possible, given the regular expression used + the
		// validation done before. Anyways, error is treated just in case
		return nil, ekerrors.New(
			moneyParseAmountError,
			fmt.Sprintf("Error converting amount '%s' in '%s' to number", amountstr, moneystr),
		)
	}

	currency, err := currency.Get(currencyCode)
	if err != nil {
		return nil, ekerrors.Wrap(
			moneyParseCurrencyError,
			err.Error(),
			err,
		)
	}

	err = shiftToCurrencyUnits(amount, currency, units)
	if err != nil {
		return nil, err
	}

	m := New(amount, currency)
	return m, nil
}

func MustParse(moneystr string) *Money {
	m, err := Parse(moneystr)
	if err != nil {
		panic(err.Error())
	}
	return m
}
