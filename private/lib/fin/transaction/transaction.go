// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/lib/fin/money"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type Transaction struct {
	slang.EKValue
}

const (
	typeName     = "Transaction"
	fSender      = "sender"
	fReceiver    = "receiver"
	fValue       = "value"
	fBookingDate = "booking_date"
	fValueDate   = "value_date"
	fPurpose     = "purpose"
	fnName       = "transaction"
)

var (
	Fn = &slang.Fn{
		Name:     fnName,
		Callback: transactionFn,
	}
)

func New(
	sender *account.Account,
	receiver *account.Account,
	value *money.Money,
	bookingDate *datetime.DateTime,
	valueDate *datetime.DateTime,
	purpose starlark.String,
) *Transaction {
	if valueDate == nil {
		valueDate = bookingDate
	}
	return &Transaction{
		slang.NewEKValue(
			typeName,
			[]string{
				fSender,
				fReceiver,
				fValue,
				fBookingDate,
				fValueDate,
				fPurpose,
			},
			map[string]starlark.Value{
				fSender:      sender,
				fReceiver:    receiver,
				fValue:       value,
				fBookingDate: bookingDate,
				fValueDate:   valueDate,
				fPurpose:     purpose,
			},
			map[string]slang.PreProcessFn{
				fSender:      account.AssertAccount,
				fReceiver:    account.AssertAccount,
				fValue:       money.AssertMoney,
				fBookingDate: datetime.AssertDateTime,
				fValueDate:   datetime.AssertDateTime,
				fPurpose:     slang.AssertString,
			},
			slang.NoMaskFn,
		),
	}
}

func transactionFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var sender, receiver *account.Account
	var value *money.Money
	var bookingDate, valueDate *datetime.DateTime
	var purpose starlark.String
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		fSender, &sender,
		fReceiver, &receiver,
		fValue, &value,
		fBookingDate, &bookingDate,
		fValueDate+"?", &valueDate,
		fPurpose+"?", &purpose,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return New(sender, receiver, value, bookingDate, valueDate, purpose), nil
}

func (tx *Transaction) Sender() *account.Account {
	return account.HasAttrsMustGetAccount(tx, fSender)
}

func (tx *Transaction) Receiver() *account.Account {
	return account.HasAttrsMustGetAccount(tx, fReceiver)
}

func (tx *Transaction) Value() *money.Money {
	return money.HasAttrsMustGetMoney(tx, fValue)
}

func (tx *Transaction) BookingDate() *datetime.DateTime {
	return datetime.HasAttrsMustGetDateTime(tx, fBookingDate)
}

func (tx *Transaction) ValueDate() *datetime.DateTime {
	return datetime.HasAttrsMustGetDateTime(tx, fValueDate)
}

func (tx *Transaction) Purpose() string {
	p := slang.HasAttrsMustGetString(tx, fPurpose)
	return string(p)
}

func (tx *Transaction) Equal(otx *Transaction) bool {
	return tx == otx || (tx.Sender().Equal(otx.Sender()) &&
		tx.Receiver().Equal(otx.Receiver()) &&
		tx.Value().Equal(otx.Value()) &&
		tx.BookingDate().Equal(otx.BookingDate()) &&
		tx.ValueDate().Equal(otx.ValueDate()) &&
		tx.Purpose() == otx.Purpose())
}
