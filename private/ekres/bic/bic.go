// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"github.com/econbits/econkit/private/ekres/country"
	"go.starlark.net/starlark"
)

type BIC struct {
	eklark.EKValue
}

const (
	typeName = "BIC"
	fCode    = "code"
	fnName   = "bic"
)

var (
	BICFn = &eklark.Fn{
		Name:     fnName,
		Callback: bicFn,
	}
)

func format(value starlark.Value) starlark.Value {
	vstr, ok := starlark.AsString(value)
	if !ok {
		return value
	}
	code := strings.ToUpper(string(vstr))

	if len(code) == 8 {
		code = code + "XXX"
	}

	return starlark.String(code)
}

func Parse(code string) (*BIC, error) {
	vcode := format(starlark.String(code))
	err := AssertBICString(vcode)
	if err != nil {
		return nil, err
	}
	return &BIC{
		eklark.NewEKValue(
			typeName,
			[]string{fCode},
			map[string]starlark.Value{
				fCode: vcode,
			},
			map[string]eklark.ValidateFn{
				fCode: AssertBICString,
			},
			map[string]eklark.FormatterFn{
				fCode: format,
			},
			eklark.NoMaskFn,
		),
	}, nil
}

func MustParse(code string) *BIC {
	bic, err := Parse(code)
	if err != nil {
		panic(err.Error())
	}
	return bic
}

func bicFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var code starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fCode, &code)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return Parse(string(code))
}

func (bic *BIC) String() string {
	return bic.Code()
}

func (bic *BIC) Code() string {
	code := eklark.HasAttrsMustGetString(bic, fCode)
	return string(code)
}

func (bic *BIC) Country() *country.Country {
	code := bic.Code()
	return country.MustGet(code[4:6])
}

func (bic *BIC) InstitutionCode() string {
	code := bic.Code()
	return code[0:4]
}

func (bic *BIC) LocationCode() string {
	code := bic.Code()
	return code[6:8]
}

func (bic *BIC) BranchCode() string {
	code := bic.Code()
	return code[8:11]
}

func (bic *BIC) Equal(obic *BIC) bool {
	return bic == obic || bic.Code() == obic.Code()
}
