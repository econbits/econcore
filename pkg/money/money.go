//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"fmt"
	"strings"

	"github.com/econbits/econcore/pkg/currency"
)

// The representation of Money is key in any financial application.
// Be aware of 3 key design decisions:
//
// - It is immutable
//
// - It can only represent absolute amounts. The reason for this is that typically
// money is used in a context. For example, given 1.00 EUR, a transfer from A to B
// means we would debit A (-1.00) and credit B (+1.00). As you see there is already
// a clear meaning of the sign (+, -) based on the operation (debit/credit). If Money
// could represent negative amounts, we could specify a debit of -1.00 which effectively
// would add +1.00. To avoid such issues, we will refer to the financial terminology
// to understand in the context if the amount is to be added or subtracted.
//
// - It is not possible to get the amount as uint64, reason being that the current implementation
// might change in the future (ex: uint128)
type Money struct {
	amount   uint64
	currency currency.Currency
}

// Returns the currency for Money
func (m Money) Currency() currency.Currency {
	return m.currency
}

// Returns the amount value as string
func (m Money) AmountStr() string {
	units := int(m.currency.Units())
	amountstr := fmt.Sprintf("%d", m.amount)
	if units == 0 {
		return amountstr
	}
	if amountstr == "0" {
		// units + 1 for: 0.00
		amountstr = strings.Repeat("0", units+1)
	}
	decPoint := len(amountstr) - units
	lAmount := amountstr[0:decPoint]
	rAmount := amountstr[decPoint:]
	return lAmount + "." + rAmount
}

// String() implementation
func (m Money) String() string {
	return m.AmountStr() + m.currency.Code()
}

// Get a Counter initialized from Money
func (m Money) Counter() Counter {
	return Counter{positive: true, m: m}
}
