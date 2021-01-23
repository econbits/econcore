// Copyright (C) 2020  Germán Fuentes Capella

package ekerrors

type Format func(msg string) string

type EKError struct {
	class   *Class
	msg     string
	wrapped error
}

func New(class *Class, msg string) *EKError {
	return &EKError{class: class, msg: msg, wrapped: nil}
}

func Wrap(class *Class, err error, formatters []Format) *EKError {
	msg := err.Error()
	if formatters != nil {
		for _, format := range formatters {
			msg = format(msg)
		}
	}
	return &EKError{class: class, msg: msg, wrapped: err}
}

func (err *EKError) Error() string {
	return "[" + err.class.s + "] " + err.msg
}

func (err *EKError) Unwrap() error {
	return err.wrapped
}

func (err *EKError) HasClass(class *Class) bool {
	return err.class.s == class.s
}
