// Copyright (C) 2020  Germán Fuentes Capella

package ekerrors

import (
	"go.starlark.net/starlark"
)

type Format func(msg string) string

type EKError struct {
	class   *Class
	msg     string
	wrapped error
	cs      starlark.CallStack
}

func New(class *Class, msg string) *EKError {
	return &EKError{class: class, msg: msg, wrapped: nil, cs: nil}
}

func Wrap(class *Class, err error, formatters []Format) *EKError {
	msg := err.Error()
	if formatters != nil {
		for _, format := range formatters {
			msg = format(msg)
		}
	}
	return &EKError{class: class, msg: msg, wrapped: err, cs: nil}
}

func (err *EKError) Error() string {
	return err.class.s + ": " + err.msg
}

func (err *EKError) Unwrap() error {
	return err.wrapped
}

func (err *EKError) HasClass(class *Class) bool {
	return err.class.s == class.s
}

func (err *EKError) LinkCS(cs starlark.CallStack) {
	err.cs = cs
}

func (err *EKError) Backtrace() string {
	if err.cs == nil {
		return ""
	}
	return err.cs.String()
}
