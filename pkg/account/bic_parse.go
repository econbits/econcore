//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"fmt"
	"strings"

	"github.com/econbits/econkit/pkg/country"
)

func ParseBIC(bicstr string) (BIC, error) {
	lbic := len(bicstr)
	if lbic != 8 && lbic != 11 {
		return BIC{}, BICFormatError(fmt.Sprintf("BIC length must be 8 or 11 characters; %d found", lbic))
	}
	bicstr = strings.ToUpper(bicstr)
	cc := bicstr[4:6]
	c, err := country.Get(cc)
	if err != nil {
		return BIC{}, err
	}
	if lbic == 8 {
		bicstr = bicstr + "XXX"
	}
	return BIC{str: bicstr, c: c}, nil
}

func MustParseBIC(bicstr string) BIC {
	bic, err := ParseBIC(bicstr)
	if err != nil {
		panic(err.Error())
	}
	return bic
}
