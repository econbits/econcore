//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"fmt"
	"regexp"
	"strconv"
)

var reMoney = regexp.MustCompile(`(?i)^\s*(?P<lh>[0-9]+)(\.(?P<rh>[0-9]+)){0,1}\s*(?P<curr>[a-z]{3}){0,1}\s*$`)

func decompose(numstr string) (string, int, string) {
	groupNames := reMoney.SubexpNames()
	matches := reMoney.FindStringSubmatch(numstr)

	amountStr := ""
	gotUnits := 0
	currStr := ""
	for i, value := range matches {
		if i == 0 {
			// this is the whole str
			continue
		}
		if groupNames[i] == "lh" {
			amountStr = value + amountStr
		} else if groupNames[i] == "rh" {
			amountStr = amountStr + value
			gotUnits = len(value)
		} else if groupNames[i] == "curr" {
			currStr = value
		}
	}
	return amountStr, gotUnits, currStr
}

func checkParsedStr(numstr string, amountStr string, currStr string) error {
	if len(amountStr) == 0 {
		return AmountNotFoundError(fmt.Sprintf("Amount not found in '%s'", numstr))
	}
	if len(currStr) == 0 {
		return CurrencyNotFoundError(fmt.Sprintf("Currency not found in '%s'", numstr))
	}
	return nil
}

func getUnitShift(m Money, gotUnits int) (uint64, error) {
	gotUnits8 := uint8(gotUnits)
	if m.currency.Units() < gotUnits8 {
		return 0, TooManyUnitsError(
			fmt.Sprintf("'%d' are too many decimal points for [%s:%d]", gotUnits, m.currency.Code(), m.currency.Units()))
	}

	shift := uint64(1)
	for ; gotUnits8 < m.currency.Units(); gotUnits8++ {
		shift = shift * 10
	}
	return shift, nil
}

func Parse(numstr string) (Money, error) {
	amountStr, gotUnits, currStr := decompose(numstr)
	err := checkParsedStr(numstr, amountStr, currStr)
	if err != nil {
		return noMoney, err
	}

	amount, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		return noMoney, AmountOverflowError(err.Error())
	}

	m, err := New(amount, currStr)
	if err != nil {
		return noMoney, err
	}

	shift, err := getUnitShift(m, gotUnits)
	if err != nil {
		return noMoney, err
	}
	m.amount = m.amount * shift
	return m, nil
}

func MustParse(numstr string) Money {
	m, err := Parse(numstr)
	if err != nil {
		panic(err.Error())
	}
	return m
}
