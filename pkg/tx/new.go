//Copyright (C) 2020  Germán Fuentes Capella

package tx

import (
	"fmt"
	"time"

	"github.com/econbits/econkit/pkg/money"
)

func New(src fmt.Stringer, dst fmt.Stringer, value money.Money, valueDate time.Time, purpose string) Tx {
	return NewWithBookingDate(src, dst, value, valueDate, valueDate, purpose)
}

func NewWithBookingDate(
	src fmt.Stringer,
	dst fmt.Stringer,
	value money.Money,
	valueDate time.Time,
	bookingDate time.Time,
	purpose string,
) Tx {
	return Tx{src: src, dst: dst, value: value, valueDate: valueDate, bookingDate: bookingDate, purpose: purpose}
}
