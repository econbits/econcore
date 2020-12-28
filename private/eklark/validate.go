// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"

	"go.starlark.net/starlark"
)

type ValidateFn func(starlark.Value) error

func IsString(v starlark.Value) error {
	_, ok := v.(starlark.String)
	if !ok {
		return fmt.Errorf("'%v' is not a string", v)
	}
	return nil
}

func IsInt(v starlark.Value) error {
	_, ok := v.(starlark.Int)
	if !ok {
		return fmt.Errorf("'%v' is not an int", v)
	}
	return nil
}

// TODO move this function to the date package
//func IsDateValue(v starlark.Value) error {
//	_, ok := v.(*Date)
//	if !ok {
//		return fmt.Errorf("'%v' is not a date", v)
//	}
//	return nil
//}
