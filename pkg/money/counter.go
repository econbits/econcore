//Copyright (C) 2020  Germán Fuentes Capella

package money

import (
	"fmt"
)

type Counter struct {
	positive bool
	m        Money
}

// true if Counter is positive
func (c Counter) IsPositive() bool {
	return c.positive
}

// true if Counter is negative
func (c Counter) IsNegative() bool {
	return !c.positive
}

// String() implementation
func (c Counter) String() string {
	sign := ""
	if c.m.amount != 0 {
		if c.positive {
			sign = "+"
		} else {
			sign = "-"
		}
	}
	return sign + c.m.String()
}

func (c Counter) Add(m Money) (Counter, error) {
	if !c.m.currency.IsEqual(m.currency) {
		errMsg := fmt.Sprintf("Counter currency: %s, Money currency: %s", c.m.currency.Code(), m.currency.Code())
		return Counter{}, CurrencyMismatchError(errMsg)
	}
	if c.positive {
		maxUint64 := ^uint64(0)
		if maxUint64-c.m.amount < m.amount {
			return Counter{}, AmountOverflowError("Counter amount overflow")
		}
		c.m.amount += m.amount
		return c, nil
	}
	// is negative
	if c.m.amount > m.amount {
		c.m.amount -= m.amount
		return c, nil
	}
	c.positive = true
	c.m.amount = m.amount - c.m.amount
	return c, nil
}

func (c Counter) Sub(m Money) (Counter, error) {
	if !c.m.currency.IsEqual(m.currency) {
		errMsg := fmt.Sprintf("Counter currency: %s, Money currency: %s", c.m.currency.Code(), m.currency.Code())
		return Counter{}, CurrencyMismatchError(errMsg)
	}
	if c.positive {
		if c.m.amount >= m.amount {
			c.m.amount -= m.amount
			return c, nil
		}
		c.positive = false
		c.m.amount = m.amount - c.m.amount
		return c, nil
	}
	// is negative
	maxUint64 := ^uint64(0)
	if maxUint64-c.m.amount < m.amount {
		return Counter{}, AmountOverflowError("Counter amount overflow")
	}
	c.m.amount += m.amount
	return c, nil
}
