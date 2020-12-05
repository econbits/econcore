//Copyright (C) 2020  Germán Fuentes Capella

package tx

import (
	"fmt"
	"time"

	"github.com/econbits/econkit/pkg/money"
)

type Tx struct {
	src         fmt.Stringer
	dst         fmt.Stringer
	value       money.Money
	bookingDate time.Time
	valueDate   time.Time
	purpose     string
}

func (tx Tx) Source() fmt.Stringer {
	return tx.src
}

func (tx Tx) Destination() fmt.Stringer {
	return tx.dst
}

func (tx Tx) Value() money.Money {
	return tx.value
}

func (tx Tx) BookingDate() time.Time {
	return tx.bookingDate
}

func (tx Tx) ValueDate() time.Time {
	return tx.valueDate
}

func (tx Tx) Purpose() string {
	return tx.purpose
}
