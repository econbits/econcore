// Copyright (C) 2020  Germán Fuentes Capella

package datetime

import (
	"fmt"
	"time"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type DateTime struct {
	eklark.EKValue
	time time.Time
}

const (
	typeName = "DateTime"
	fLayout  = "layout"
	fValue   = "value"
	fnName   = "datetime"
)

var (
	errorClass = ekerrors.MustRegisterClass("DateTimeError")
	DateTimeFn = &eklark.Fn{
		Name:     fnName,
		Callback: dateTimeFn,
	}
)

func New(layout string, value string) (*DateTime, error) {
	t, err := time.Parse(string(layout), string(value))
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return NewFromValues(starlark.String(layout), starlark.String(value), t), nil
}

func NewFromValues(layout starlark.String, value starlark.String, t time.Time) *DateTime {
	return &DateTime{
		eklark.NewEKValue(
			typeName,
			[]string{fLayout, fValue},
			map[string]starlark.Value{
				fLayout: layout,
				fValue:  value,
			},
			map[string]eklark.ValidateFn{
				fLayout: eklark.AssertString,
				fValue:  eklark.AssertString,
			},
			map[string]eklark.FormatterFn{},
			eklark.NoMaskFn,
		),
		t,
	}
}

func NewFromTime(t time.Time) *DateTime {
	dt := NewFromValues(
		starlark.String(time.RFC3339),
		starlark.String(t.Format(time.RFC3339)),
		t,
	)
	return dt
}

func dateTimeFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var layout, value starlark.String
	err := starlark.UnpackArgs(builtin.Name(), args, kwargs, fLayout, &layout, fValue, &value)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err.Error(),
			err,
		)
	}
	return New(string(layout), string(value))
}

func AssertDateTime(v starlark.Value) error {
	_, ok := v.(*DateTime)
	if !ok {
		return fmt.Errorf("'%v' is not a date", v)
	}
	return nil
}

func (dt *DateTime) Time() time.Time {
	return dt.time
}

func (dt *DateTime) String() string {
	str := eklark.HasAttrsMustGetString(dt, fValue)
	return string(str)
}
