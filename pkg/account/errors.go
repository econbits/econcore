//Copyright (C) 2020  Germán Fuentes Capella

package account

type IBANFormatError string

func (e IBANFormatError) Error() string {
	return string(e)
}

type BICFormatError string

func (e BICFormatError) Error() string {
	return string(e)
}
