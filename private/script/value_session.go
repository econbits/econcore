//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

type Session struct {
	d starlark.StringDict
}

// Implementing starlark Value interface

func (s Session) String() string {
	return fmt.Sprintf("Session{%v}", s.d)
}

func (s Session) Type() string {
	return "Session"
}

func (s Session) Freeze() {
	s.d.Freeze()
}

func (s Session) Truth() starlark.Bool {
	return len(s.d) > 0
}

func (s Session) Hash() (uint32, error) {
	return 0, fmt.Errorf("Session can't be hashed")
}

// Implementing starlark Mapping interface

func (s Session) Get(k starlark.Value) (v starlark.Value, found bool, err error) {
	ks, ok := starlark.AsString(k)
	if !ok {
		return nil, false, fmt.Errorf("Session only accepts string keys; found: '%v'", k)
	}
	v, found = s.d[ks]
	if !found {
		return nil, false, nil
	}
	return v, true, nil
}

// Implementing starlark HasSetKey interface

func (s Session) SetKey(k, v starlark.Value) error {
	ks, ok := starlark.AsString(k)
	if !ok {
		return fmt.Errorf("Session only accepts string keys; found: '%v'", k)
	}
	s.d[ks] = v
	return nil
}
