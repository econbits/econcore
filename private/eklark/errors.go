// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"
)

type ErrorType string

func (et ErrorType) String() string {
	return string(et)
}

type EKError struct {
	FilePath    string
	Function    string
	ErrorType   ErrorType
	Description string
}

func (eke *EKError) Error() string {
	return fmt.Sprintf("[%s][%s] %s", ScriptId(eke.FilePath), eke.Function, eke.Description)
}
