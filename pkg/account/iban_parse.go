//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"regexp"
	"strings"

	"github.com/econbits/econkit/pkg/country"
)

var (
	reIBAN = regexp.MustCompile(`[A-Za-z]{2}[0-9]{2}[0-9A-Za-z]{4,30}`)
)

func ParseIBAN(ibanstr string) (IBAN, error) {
	// electronic form IBAN
	efIban := strings.ReplaceAll(ibanstr, " ", "")
	if !reIBAN.MatchString(efIban) {
		return IBAN{}, IBANFormatError("'" + ibanstr + "' is not a valid IBAN")
	}
	country, err := country.Get(efIban[0:2])
	if err != nil {
		return IBAN{}, err
	}
	return IBAN{n: efIban, c: country}, nil
}

func MustParseIBAN(ibanstr string) IBAN {
	iban, err := ParseIBAN(ibanstr)
	if err != nil {
		panic(err.Error())
	}
	return iban
}

var (
	br = country.MustGet("BR")
	mu = country.MustGet("MU")
)

func form4(ibanstr string, ibanlen int) string {
	i, pf := 0, ""
	for ; i < ibanlen/4; i++ {
		if len(pf) > 0 {
			pf = pf + " "
		}
		iln := 4 * i
		pf = pf + ibanstr[iln:iln+4]
	}
	iln := i * 4
	if ibanlen > iln {
		pf = pf + " " + ibanstr[iln:ibanlen]
	}
	return pf
}

// Printed Form String (in groups of 4 characters)
func (iban IBAN) PrintedForm() string {
	country := iban.Country()
	if country.IsEqual(br) && len(iban.n) == 29 {
		ibanstr := form4(iban.n, 24)
		return ibanstr + " " + iban.n[24:27] + " " + iban.n[27:29]
	}
	if country.IsEqual(mu) && len(iban.n) == 30 {
		ibanstr := form4(iban.n, 24)
		return ibanstr + " " + iban.n[24:27] + " " + iban.n[27:30]
	}
	return form4(iban.n, len(iban.n))
}

func (iban IBAN) ElectronicForm() string {
	return iban.String()
}
