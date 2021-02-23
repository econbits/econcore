// Copyright (C) 2021  Germán Fuentes Capella

package json

import (
	pjson "encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	decFnNameme = "json_decode"
)

var (
	DecFn = &slang.Fn{
		Name:     decFnNameme,
		Callback: decodeFn,
	}
)

func parse(i interface{}) (starlark.Value, error) {
	if i == nil {
		return starlark.None, nil
	}
	switch itype := i.(type) {
	case float64:
		return starlark.Float(itype), nil
	case string:
		return starlark.String(itype), nil
	case bool:
		return starlark.Bool(itype), nil
	default:
		val := reflect.ValueOf(i)
		if val.Kind() == reflect.Map {
			jd := starlark.NewDict(10)
			for _, k := range val.MapKeys() {
				v := val.MapIndex(k)
				kvalue, err := parse(k.Interface())
				if err != nil {
					return nil, err
				}
				vvalue, err := parse(v.Interface())
				if err != nil {
					return nil, err
				}
				err = jd.SetKey(kvalue, vvalue)
				if err != nil {
					return nil, err
				}
			}
			return jd, nil
		}
		if val.Kind() == reflect.Slice {
			jl := starlark.NewList([]starlark.Value{})
			for i := 0; i < val.Len(); i++ {
				obj := val.Index(i)
				ivalue, err := parse(obj.Interface())
				if err != nil {
					return nil, err
				}
				err = jl.Append(ivalue)
				if err != nil {
					return nil, err
				}
			}
			return jl, nil
		}
		return nil, errors.New(fmt.Sprintf("unknown json type: %T", i))
	}
}

func decodeFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var text starlark.String
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		"text", &text,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	if text == starlark.String("") {
		return text, nil
	}
	var j interface{}
	err = pjson.Unmarshal([]byte(text), &j)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	jobj, err := parse(j)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	return jobj, nil
}
