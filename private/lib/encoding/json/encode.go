// Copyright (C) 2021  Germán Fuentes Capella

package json

import (
	pjson "encoding/json"
	"errors"
	"fmt"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

const (
	encFnNameme = "json_encode"
)

var (
	EncFn = &slang.Fn{
		Name:     encFnNameme,
		Callback: encodeFn,
	}
)

func convert(v starlark.Value) (interface{}, error) {
	if v == starlark.None {
		return nil, nil
	}
	switch vtype := v.(type) {
	case starlark.Float:
		return float64(vtype), nil
	case starlark.String:
		return string(vtype), nil
	case starlark.Bool:
		return bool(vtype), nil
	case *starlark.Dict:
		d := map[string]interface{}{}
		for _, t := range vtype.Items() {
			kt, ok := t[0].(starlark.String)
			if !ok {
				return nil, errors.New(fmt.Sprintf("key %v should be of type string", t[0]))
			}
			vt, err := convert(t[1])
			if err != nil {
				return nil, err
			}
			d[string(kt)] = vt
		}
		return d, nil
	case *starlark.List:
		l := []interface{}{}
		for i := 0; i < vtype.Len(); i++ {
			elem, err := convert(vtype.Index(i))
			if err != nil {
				return nil, err
			}
			l = append(l, elem)
		}
		return l, nil
	default:
		return nil, errors.New(fmt.Sprintf("can't convert %v to a json valid type", vtype.Type()))
	}
}

func encodeFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var obj starlark.Value
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		"obj", &obj,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	gobj, err := convert(obj)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	jbytes, err := pjson.Marshal(gobj)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	return starlark.String(jbytes), nil
}
