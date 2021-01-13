//Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"testing"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/bic/"
	fn := BICFn
	epilogue := starlark.StringDict{fn.Name: fn.Builtin()}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}

func TestParse(t *testing.T) {
	codes := []string{SampleDE, "deutdeffXXX", "DEUTDEFF"}
	for _, code := range codes {
		bic := MustParse(code)
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
	codes := []string{"DEUT", "deutdeff500000"}
	for _, code := range codes {
		_, err := Parse(code)
		if err == nil {
			t.Errorf("Expected error; got none")
		}
		if len(err.Error()) == 0 {
			t.Errorf("Expected error message; none found")
		}
	}
}

func TestParseBICWrongCountry(t *testing.T) {
	code := "DEUTZZFF500"
	_, err := Parse(code)
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
			t.Errorf("MustParse did not panic")
		}
	}()

	code := "DEUT"
	MustParse(code)
}

func TestEqual(t *testing.T) {
	bic1 := MustParse(SampleDE)
	bic2 := MustParse(SampleZA)

	if !bic1.Equal(bic1) {
		t.Fatalf("%v should be equal to %v", bic1, bic1)
	}

	if bic1.Equal(bic2) {
		t.Fatalf("%v should not be equal to %v", bic1, bic2)
	}
}
