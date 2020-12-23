//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"go.starlark.net/starlark"
)

const (
	sessionType = ekmValueType("Session")
)

type Session struct {
	EKMValue
}

// New function

func NewSession() *Session {
	return &Session{
		EKMValue{
			valueType:    sessionType,
			attrs:        []string{},
			data:         map[string]starlark.Value{},
			validatorsFn: map[string]validatorFunc{},
			frozen:       false,
			mask:         noMask,
		},
	}
}
