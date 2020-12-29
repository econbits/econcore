// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"

	"go.starlark.net/starlark"
)

func call(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	fn := &Fn{
		Name:     "TestFn",
		ArgNames: []string{"arg1"},
		Callback: func(
			thread *starlark.Thread,
			builtin *starlark.Builtin,
			sdict StringDict,
		) (starlark.Value, error) {
			return sdict["arg1"], nil
		},
		ArgError: ErrorType("FnError"),
	}
	thread := NewThread("TestThread")
	return fn.Builtin().CallInternal(thread, args, kwargs)
}

func TestFnCallbackSuccess(t *testing.T) {
	value, err := call(starlark.Tuple{starlark.True}, []starlark.Tuple{})
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	tvalue, ok := value.(starlark.Bool)
	if !ok {
		t.Fatalf("value %v is not a bool", value)
	}
	if !tvalue.Truth() {
		t.Fatalf("value %v is not a bool", value)
	}
}

func TestFnCallbackSuccessKWArgs(t *testing.T) {
	value, err := call(
		starlark.Tuple{},
		[]starlark.Tuple{[]starlark.Value{starlark.String("arg1"), starlark.True}},
	)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	tvalue, ok := value.(starlark.Bool)
	if !ok {
		t.Fatalf("value %v is not a bool", value)
	}
	if !tvalue.Truth() {
		t.Fatalf("value %v is not a bool", value)
	}
}

func TestFnCallbackKWArgsWithWrongKeyType(t *testing.T) {
	_, err := call(
		starlark.Tuple{},
		[]starlark.Tuple{[]starlark.Value{starlark.MakeInt(1), starlark.True}},
	)
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestFnCallbackKWArgsWithWrongKeyName(t *testing.T) {
	_, err := call(
		starlark.Tuple{},
		[]starlark.Tuple{[]starlark.Value{starlark.String("wrong"), starlark.True}},
	)
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestFnCallbackEmptyArgs(t *testing.T) {
	value, err := call(
		starlark.Tuple{},
		[]starlark.Tuple{},
	)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	_, ok := value.(starlark.NoneType)
	if !ok {
		t.Fatalf("value %v is not None", value)
	}
	if value.Truth() {
		t.Fatalf("value %v is not false", value)
	}
}

func TestFnCallbackTooManyArgs(t *testing.T) {
	_, err := call(
		starlark.Tuple{starlark.True, starlark.True},
		[]starlark.Tuple{},
	)
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}
