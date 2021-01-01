// Copyright (C) 2020  Germán Fuentes Capella

package date

import (
	"fmt"
	"time"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type Date struct {
	eklark.EKValue
}

const (
	dateType    = "Date"
	fDateLayout = "layout"
	fDateValue  = "value"
	fnName      = "date"
)

var (
	DateError = ekerrors.MustRegisterClass("DateError")
	DateFn    = &eklark.Fn{
		Name:     fnName,
		Callback: dateFn,
	}
)

func New(layout string, value string) *Date {
	return NewFromValues(starlark.String(layout), starlark.String(value))
}

func NewFromValues(layout starlark.String, value starlark.String) *Date {
	return &Date{
		eklark.NewEKValue(
			dateType,
			[]string{fDateLayout, fDateValue},
			map[string]starlark.Value{
				fDateLayout: layout,
				fDateValue:  value,
			},
			map[string]eklark.ValidateFn{
				fDateLayout: eklark.AssertString,
				fDateValue:  eklark.AssertString,
			},
			eklark.NoMaskFn,
		),
	}
}

func NewFromTime(t time.Time) *Date {
	return New(time.RFC3339, t.Format(time.RFC3339))
}

func dateFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var layout, value starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fDateLayout, &layout, fDateValue, &value)
	if err != nil {
		return nil, ekerrors.Wrap(
			DateError,
			err.Error(),
			err,
		)
	}
	_, err = time.Parse(string(layout), string(value))
	if err != nil {
		return nil, ekerrors.Wrap(
			DateError,
			err.Error(),
			err,
		)
	}
	return NewFromValues(layout, value), nil
}

func AssertDate(v starlark.Value) error {
	_, ok := v.(*Date)
	if !ok {
		return fmt.Errorf("'%v' is not a date", v)
	}
	return nil
}

func (datev *Date) Time() (time.Time, error) {
	layout := eklark.HasAttrsMustGetString(datev, fDateLayout)
	value := eklark.HasAttrsMustGetString(datev, fDateValue)
	return time.Parse(string(layout), string(value))
}
