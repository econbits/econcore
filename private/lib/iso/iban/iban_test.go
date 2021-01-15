//Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"strings"
	"testing"
)

var (
	ibans = []string{
		SampleBE,
		SampleBR,
		SampleFR,
		SampleDE,
		SampleGR,
		SampleMU,
		SamplePK,
		SamplePL,
		SampleRO,
		SampleLC,
		SampleSA,
		SampleES,
		SampleCH,
		SampleGB,
	}
)

func TestParseIBAN(t *testing.T) {
	for _, ibanstr := range ibans {
		iban, err := Parse(ibanstr)
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
		MustParse(ibanstr)
	}
}

func TestParseIBANWrongPattern(t *testing.T) {
	ibanstr := "this is not an iban"
	_, err := Parse(ibanstr)
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}

func TestParseIBANWrongCountry(t *testing.T) {
	ibanstr := "ZZ71 0961 2345 6769"
	_, err := Parse(ibanstr)
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}

func TestMustParseIBANWrongPattern(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseIBAN did not panic")
		}
	}()

	ibanstr := "this is not an iban"
	MustParse(ibanstr)
}
