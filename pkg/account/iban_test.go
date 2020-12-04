//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"strings"
	"testing"
)

// These IBANs are taken from:
// https://en.wikipedia.org/wiki/International_Bank_Account_Number
const (
	ibanBE = "BE71 0961 2345 6769"
	ibanBR = "BR15 0000 0000 0000 1093 2840 814 P2"
	ibanFR = "FR76 3000 6000 0112 3456 7890 189"
	ibanDE = "DE91 1000 0000 0123 4567 89"
	ibanGR = "GR96 0810 0010 0000 0123 4567 890"
	ibanMU = "MU43 BOMM 0101 1234 5678 9101 000 MUR"
	ibanPK = "PK70 BANK 0000 1234 5678 9000"
	ibanPL = "PL10 1050 0099 7603 1234 5678 9123"
	ibanRO = "RO09 BCYP 0000 0012 3456 7890"
	ibanLC = "LC14 BOSL 1234 5678 9012 3456 7890 1234"
	ibanSA = "SA44 2000 0001 2345 6789 1234"
	ibanES = "ES79 2100 0813 6101 2345 6789"
	ibanCH = "CH56 0483 5012 3456 7800 9"
	ibanGB = "GB98 MIDL 0700 9312 3456 78"
)

var (
	ibans = []string{
		ibanBE,
		ibanBR,
		ibanFR,
		ibanDE,
		ibanGR,
		ibanMU,
		ibanPK,
		ibanPL,
		ibanRO,
		ibanLC,
		ibanSA,
		ibanES,
		ibanCH,
		ibanGB,
	}
)

func TestParseIBAN(t *testing.T) {
	for _, ibanstr := range ibans {
		iban, err := ParseIBAN(ibanstr)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		} else {
			efIban := strings.ReplaceAll(ibanstr, " ", "")
			if efIban != iban.String() {
				t.Errorf("Expected : %s; got: %s", ibanstr, efIban)
			}
			if efIban != iban.ElectronicForm() {
				t.Errorf("Expected : %s; got: %s", ibanstr, efIban)
			}
			pfIban := iban.PrintedForm()
			if pfIban != ibanstr {
				t.Errorf("Expected : %s; got: %s", ibanstr, pfIban)
			}
		}
	}
}

func TestParseIBANWrongPattern(t *testing.T) {
	ibanstr := "this is not an iban"
	_, err := ParseIBAN(ibanstr)
	if err == nil {
		t.Errorf("Expected error; none found")
	}
	if len(err.Error()) == 0 {
		t.Errorf("Expected error message; none found")
	}
}

func TestParseIBANWrongCountry(t *testing.T) {
	ibanstr := "ZZ71 0961 2345 6769"
	_, err := ParseIBAN(ibanstr)
	if err == nil {
		t.Errorf("Expected error; none found")
	}
	if len(err.Error()) == 0 {
		t.Errorf("Expected error message; none found")
	}
}

func TestMustParseIBANWrongPattern(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseIBAN did not panic")
		}
	}()

	ibanstr := "this is not an iban"
	MustParseIBAN(ibanstr)
}
