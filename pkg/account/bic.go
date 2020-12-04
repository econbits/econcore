//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"github.com/econbits/econkit/pkg/country"
)

// BIC based on https://en.wikipedia.org/wiki/ISO_9362
type BIC struct {
	str string
	c   country.Country
}

func (bic BIC) InstitutionCode() string {
	return bic.str[0:4]
}

func (bic BIC) Country() country.Country {
	return bic.c
}

func (bic BIC) LocationCode() string {
	return bic.str[6:8]
}

func (bic BIC) BranchCode() string {
	return bic.str[8:11]
}

func (bic BIC) String() string {
	return bic.str
}
