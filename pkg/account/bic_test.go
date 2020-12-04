//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"testing"
)

func TestParseBIC(t *testing.T) {
	bicstrs := []string{"DEUTDEFFXXX", "deutdeffXXX", "DEUTDEFF"}
	for _, bicstr := range bicstrs {
		bic, err := ParseBIC(bicstr)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
		if bic.InstitutionCode() != "DEUT" {
			t.Errorf("Expected Institution Code: 'DEUT'; got: '%s'", bic.InstitutionCode())
		}
		if bic.Country().Alpha2() != "DE" {
			t.Errorf("Expected Country Code: 'DE'; got: '%s'", bic.Country().Alpha2())
		}
		if bic.LocationCode() != "FF" {
			t.Errorf("Expected Location Code: 'FF'; got: '%s'", bic.LocationCode())
		}
		if bic.BranchCode() != "XXX" {
			t.Errorf("Expected Branch Code: 'XXX'; got: '%s'", bic.BranchCode())
		}
		if bic.String() != "DEUTDEFFXXX" {
			t.Errorf("Expected BIC: 'DEUTDEFFXXX'; got: '%s'", bic.String())
		}
	}
}

func TestParseBICWrongLength(t *testing.T) {
	bicstrs := []string{"DEUT", "deutdeff500000"}
	for _, bicstr := range bicstrs {
		_, err := ParseBIC(bicstr)
		if err == nil {
			t.Errorf("Expected error; got none")
		}
		if len(err.Error()) == 0 {
			t.Errorf("Expected error message; none found")
		}
	}
}

func TestParseBICWrongCountry(t *testing.T) {
	bicstr := "DEUTZZFF500"
	_, err := ParseBIC(bicstr)
	if err == nil {
		t.Errorf("Expected error; got none")
	}
	if len(err.Error()) == 0 {
		t.Errorf("Expected error message; none found")
	}
}

func TestMustParseBICWrongLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseBIC did not panic")
		}
	}()

	bicstr := "DEUT"
	MustParseBIC(bicstr)
}
