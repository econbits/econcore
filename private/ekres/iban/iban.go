// Copyright (C) 2021  Germán Fuentes Capella

package iban

import (
	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/country"
	"go.starlark.net/starlark"
)

type IBAN struct {
	eklark.EKValue
}

const (
	typeName = "IBAN"
	fNumber  = "number"
	fnName   = "iban"
)

var (
	IBANFn = &eklark.Fn{
		Name:     fnName,
		Callback: ibanFn,
	}

	countryBR = country.MustGet("BR")
	countryMU = country.MustGet("MU")
)

func Parse(number string) (*IBAN, error) {
	vnumber, err := preprocess(starlark.String(number))
	if err != nil {
		return nil, err
	}
	return &IBAN{
		eklark.NewEKValue(
			typeName,
			[]string{fNumber},
			map[string]starlark.Value{
				fNumber: vnumber,
			},
			map[string]eklark.PreProcessFn{
				fNumber: preprocess,
			},
			eklark.NoMaskFn,
		),
	}, nil
}

func MustParse(number string) *IBAN {
	iban, err := Parse(number)
	if err != nil {
		panic(err.Error())
	}
	return iban
}

func ibanFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var number starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fNumber, &number)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return Parse(string(number))
}

func (iban *IBAN) String() string {
	return iban.ElectronicForm()
}

func (iban *IBAN) Country() *country.Country {
	number := iban.ElectronicForm()
	return country.MustGet(number[0:2])
}

func (iban *IBAN) Equal(oiban *IBAN) bool {
	return iban == oiban || iban.ElectronicForm() == oiban.ElectronicForm()
}

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
func (iban *IBAN) PrintedForm() string {
	ibanstr := iban.ElectronicForm()
	liban := len(ibanstr)
	country := iban.Country()
	if country.Equal(countryBR) && liban == 29 {
		return form4(ibanstr, 24) + " " + ibanstr[24:27] + " " + ibanstr[27:29]
	}
	if country.Equal(countryMU) && liban == 30 {
		return form4(ibanstr, 24) + " " + ibanstr[24:27] + " " + ibanstr[27:30]
	}
	return form4(ibanstr, liban)
}

func (iban *IBAN) ElectronicForm() string {
	number := eklark.HasAttrsMustGetString(iban, fNumber)
	return string(number)
}
