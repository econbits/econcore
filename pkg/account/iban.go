//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/pkg/country"
)

type IBAN struct {
	n string
	c country.Country
}

func (i IBAN) String() string {
	return i.n
}

func (i IBAN) Country() country.Country {
	return i.c
}
