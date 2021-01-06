// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"

	"go.starlark.net/starlark"
)

func getTestHasAttrsValue(attrname string, attrvalue starlark.Value) starlark.HasAttrs {
	type_ := "TestValue"
	tv := &TestValue{
		NewEKValue(
			type_,
			[]string{attrname},
			map[string]starlark.Value{
				attrname: attrvalue,
			},
			map[string]PreProcessFn{},
			NoMaskFn,
		),
	}
	return tv
}

func TestHasAttrsMustGetString(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")
	tv := getTestHasAttrsValue(attrname, attrvalue)

	gotvalue := HasAttrsMustGetString(tv, attrname)
	if gotvalue != attrvalue {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}

	gotvalue2 := HasAttrsMustGet(tv, attrname)
	gotvalue, ok := gotvalue2.(starlark.String)
	if !ok {
		t.Fatalf("expected string; found '%T'", gotvalue2)
	}
	if gotvalue != attrvalue {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetStringWithInvalidType(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.MakeInt(1)
	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetString(tv, attrname)
}

func TestHasAttrsMustGetStringWithInvalidAttrName(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")
	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetString(tv, "this attr name does not exist")
}

func TestHasAttrsMustGetWithInvalidAttrName(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")
	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGet(tv, "this attr name does not exist")
}

func TestHasAttrsMustGetInt(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.MakeInt(1)
	tv := getTestHasAttrsValue(attrname, attrvalue)

	gotvalue := HasAttrsMustGetInt(tv, attrname)
	if gotvalue != attrvalue {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetIntWithInvalidType(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")
	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetInt(tv, attrname)
}

func TestHasAttrsMustGetIntWithInvalidAttrName(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.MakeInt(1)
	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetInt(tv, "this attr name does not exist")
}

func TestHasAttrsMustGetGoInt(t *testing.T) {
	attrname := "attr"
	attrvalue := 1
	tv := getTestHasAttrsValue(attrname, starlark.MakeInt(attrvalue))

	gotvalue := HasAttrsMustGetGoInt(tv, attrname)
	if gotvalue != attrvalue {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetGoIntTooBig(t *testing.T) {
	attrname := "attr"
	attrvalue := int64(^uint64(0) >> 1)
	tv := getTestHasAttrsValue(attrname, starlark.MakeInt64(attrvalue))

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetGoInt(tv, attrname)
}
