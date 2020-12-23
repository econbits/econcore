//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"go.starlark.net/starlark"
)

type Transaction struct {
	EKMValue
}

const (
	transactionType = ekmValueType("Transaction")
	txSrcIban       = "src_iban"
	txSrcBic        = "src_bic"
	txDstIban       = "dst_iban"
	txDstBic        = "dst_bic"
	txAmount        = "amount"
	txCurrency      = "currency"
	txBookingDate   = "booking_date"
	txValueDate     = "value_date"
	txPurpose       = "purpose"
)

// New function

func NewTransaction() *Transaction {
	return &Transaction{
		EKMValue{
			valueType: transactionType,
			attrs: []string{
				txSrcIban,
				txSrcBic,
				txDstIban,
				txDstBic,
				txAmount,
				txCurrency,
				txBookingDate,
				txValueDate,
				txPurpose,
			},
			data: map[string]starlark.Value{},
			validatorsFn: map[string]validatorFunc{
				credUsername:  isStringValue,
				txSrcIban:     isStringValue,
				txSrcBic:      isStringValue,
				txDstIban:     isStringValue,
				txDstBic:      isStringValue,
				txAmount:      isIntValue,
				txCurrency:    isStringValue,
				txBookingDate: isStringValue,
				txValueDate:   isStringValue,
				txPurpose:     isStringValue,
			},
			frozen: false,
			mask:   noMask,
		},
	}
}
