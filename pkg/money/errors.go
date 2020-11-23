//Copyright (C) 2020  Germán Fuentes Capella

package money

type AmountNotFoundError string

func (e AmountNotFoundError) Error() string {
	return string(e)
}

type CurrencyNotFoundError string

func (e CurrencyNotFoundError) Error() string {
	return string(e)
}

type TooManyUnitsError string

func (e TooManyUnitsError) Error() string {
	return string(e)
}

type AmountOverflowError string

func (e AmountOverflowError) Error() string {
	return string(e)
}
