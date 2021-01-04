// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"

	"go.starlark.net/starlark"
)

type PreProcessFn func(starlark.Value) (starlark.Value, error)

func AssertString(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(starlark.String)
	if !ok {
		return nil, fmt.Errorf("'%v' is not a string", v)
	}
	return v, nil
}

func AssertInt(v starlark.Value) (starlark.Value, error) {
	_, ok := v.(starlark.Int)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an int", v)
	}
	return v, nil
}

func AssertInt32(v starlark.Value) (starlark.Value, error) {
	intv, ok := v.(starlark.Int)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an int", v)
	}
	_, err := starlark.AsInt32(intv)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func AssertInt64(v starlark.Value) (starlark.Value, error) {
	intv, ok := v.(starlark.Int)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an int", v)
	}
	_, ok = intv.Int64()
	if !ok {
		return nil, fmt.Errorf("'%v' is not a 64 bits int", v)
	}
	return v, nil
}

func AssertUint64(v starlark.Value) (starlark.Value, error) {
	intv, ok := v.(starlark.Int)
	if !ok {
		return nil, fmt.Errorf("'%v' is not an int", v)
	}
	_, ok = intv.Uint64()
	if !ok {
		return nil, fmt.Errorf("'%v' is not an unsigned 64 bits int", v)
	}
	return v, nil
}
