// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"

	"go.starlark.net/starlark"
)

type ValidateFn func(starlark.Value) error

func AssertString(v starlark.Value) error {
	_, ok := v.(starlark.String)
	if !ok {
		return fmt.Errorf("'%v' is not a string", v)
	}
	return nil
}

func AssertInt(v starlark.Value) error {
	_, ok := v.(starlark.Int)
	if !ok {
		return fmt.Errorf("'%v' is not an int", v)
	}
	return nil
}

func AssertInt32(v starlark.Value) error {
	intv, ok := v.(starlark.Int)
	if !ok {
		return fmt.Errorf("'%v' is not an int", v)
	}
	_, err := starlark.AsInt32(intv)
	return err
}

func AssertInt64(v starlark.Value) error {
	intv, ok := v.(starlark.Int)
	if !ok {
		return fmt.Errorf("'%v' is not an int", v)
	}
	_, ok = intv.Int64()
	if !ok {
		return fmt.Errorf("'%v' is not a 64 bits int", v)
	}
	return nil
}

func AssertUint64(v starlark.Value) error {
	intv, ok := v.(starlark.Int)
	if !ok {
		return fmt.Errorf("'%v' is not an int", v)
	}
	_, ok = intv.Uint64()
	if !ok {
		return fmt.Errorf("'%v' is not an unsigned 64 bits int", v)
	}
	return nil
}
