//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"
)

type ErrorType uint32

const (
	unknownError         = ErrorType(0) // do not include in errorTypes, catch all error
	LoadError            = ErrorType(1)
	ReservedVarError     = ErrorType(2)
	MissingFunctionError = ErrorType(3)
	FunctionCallError    = ErrorType(4)
	SessionError         = ErrorType(5)
	AccountListError     = ErrorType(6)
	AccountError         = ErrorType(7)
	TransactionListError = ErrorType(8)
	TransactionError     = ErrorType(9)
)

var (
	errorTypes = []ErrorType{
		LoadError,
		ReservedVarError,
		MissingFunctionError,
		FunctionCallError,
		SessionError,
		AccountListError,
		AccountError,
		TransactionListError,
		TransactionError,
	}
)

func (et ErrorType) mustTypeString() string {
	if et == LoadError {
		return "LoadError"
	}
	if et == ReservedVarError {
		return "ReservedVarError"
	}
	if et == MissingFunctionError {
		return "MissingFunctionError"
	}
	if et == FunctionCallError {
		return "FunctionCallError"
	}
	if et == SessionError {
		return "SessionError"
	}
	if et == AccountListError {
		return "AccountListError"
	}
	if et == AccountError {
		return "AccountError"
	}
	if et == TransactionListError {
		return "TransactionListError"
	}
	if et == TransactionError {
		return "TransactionError"
	}
	panic(fmt.Sprintf("ErrorType(%d) not defined", uint32(et)))
}

type ScriptError struct {
	scriptName string
	function   string
	errorType  ErrorType
	text       string
}

func (se ScriptError) Error() string {
	if len(se.scriptName) > 0 {
		return fmt.Sprintf("[%s][%s] %s", se.scriptName, se.function, se.text)
	}
	return fmt.Sprintf("{%s} %s", se.function, se.text)
}

func newLoginError(scriptName string, errorType ErrorType, text string) error {
	return ScriptError{scriptName: scriptName, function: "login", errorType: errorType, text: text}
}

func newAccountError(scriptName string, errorType ErrorType, text string) error {
	return ScriptError{scriptName: scriptName, function: "account", errorType: errorType, text: text}
}

func newBuiltinError(function string, errorType ErrorType, text string) error {
	return ScriptError{scriptName: "", function: function, errorType: errorType, text: text}
}
