// Copyright (C) 2020  Germán Fuentes Capella

package session

import (
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type Session struct {
	eklark.EKValue
}

const (
	sessionType = "Session"
	fnName      = "session"
)

var (
	sessionErrorClass = ekerrors.MustRegisterClass("SessionError")
	SessionFn         = &eklark.Fn{
		Name:     fnName,
		Callback: sessionFn,
	}
)

func New() *Session {
	return &Session{
		eklark.NewEKValue(
			sessionType,
			[]string{},
			map[string]starlark.Value{},
			map[string]eklark.ValidateFn{},
			map[string]eklark.FormatterFn{},
			eklark.NoMaskFn,
		),
	}
}

func sessionFn(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	if len(args) > 0 {
		return nil, ekerrors.New(
			sessionErrorClass,
			fmt.Sprintf("unnamed arguments are not allowed: %v", args),
		)
	}
	s := New()
	if len(kwargs) > 0 {
		for _, tuple := range kwargs {
			s.SetKey(tuple[0], tuple[1])
		}
	}
	return s, nil
}
