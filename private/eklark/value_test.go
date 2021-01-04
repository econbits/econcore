// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"
	"testing"

	"go.starlark.net/starlark"
)

type TestValue struct {
	EKValue
}

func getValueType(type_, attrname, attrvalue string) *TestValue {
	tv := &TestValue{
		NewEKValue(
			type_,
			[]string{attrname},
			map[string]starlark.Value{},
			map[string]PreProcessFn{
				attrname: AssertString,
			},
			NoMaskFn,
		),
	}
	if len(attrname) > 0 {
		tv.SetField(attrname, starlark.String(attrvalue))
	}
	return tv
}

func getMapValueType(type_, attrname, attrvalue string) *TestValue {
	tv := &TestValue{
		NewEKValue(
			type_,
			[]string{},
			map[string]starlark.Value{
				attrname: starlark.String(attrvalue),
			},
			map[string]PreProcessFn{
				attrname: AssertString,
			},
			NoMaskFn,
		),
	}
	return tv
}

func TestValueType(t *testing.T) {
	type_ := "TestValue"
	tv := getValueType(type_, "attr", "value")
	if tv.Type() != type_ {
		t.Fatalf("Expected '%s'; got '%s'", type_, tv.Type())
	}
}

func TestValueGetAttr(t *testing.T) {
	type_ := "TestValue"
	attrname := "attr"
	attrvalue := "value"
	tv := getValueType(type_, attrname, attrvalue)
	value, err := tv.Attr(attrname)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	strvalue, ok := starlark.AsString(value)
	if !ok {
		t.Fatalf("Expected string; found '%T'", value)
	}
	if strvalue != attrvalue {
		t.Fatalf("Expected '%s'; got '%s'", attrvalue, strvalue)
	}

	value, found, err := tv.Get(starlark.String(attrname))
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if !found {
		t.Fatal("Expected attr to be found; missing")
	}

	strvalue, ok = starlark.AsString(value)
	if !ok {
		t.Fatalf("Expected string; found '%T'", value)
	}
	if strvalue != attrvalue {
		t.Fatalf("Expected '%s'; got '%s'", attrvalue, strvalue)
	}
}

func TestValueGetInvalidAttrName(t *testing.T) {
	type_ := "TestValue"
	attrname := "attr"
	attrvalue := "value"
	tv := getValueType(type_, attrname, attrvalue)
	_, err := tv.Attr("this attr name does not exist")
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestValueGetDynamicAttrName(t *testing.T) {
	type_ := "TestValue"
	tv := getMapValueType(type_, "", "")
	value, err := tv.Attr("attrname")
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if value != nil {
		t.Fatalf("Expected 'nil'; got '%v'", value)
	}
}

func TestValueString(t *testing.T) {
	type_ := "TestValue"
	attrname := "attr"
	attrvalue := "value"
	tv := getValueType(type_, attrname, attrvalue)
	gotstr := tv.String()

	estr := fmt.Sprintf("%s{%s=\"%s\"}", type_, attrname, attrvalue)
	if gotstr != estr {
		t.Fatalf("Expected '%s'; got '%s'", estr, gotstr)
	}
}

func TestValueTruth(t *testing.T) {
	type_ := "TestValue"
	attrname := "attr"
	attrvalue := "value"
	tv := getValueType(type_, attrname, attrvalue)

	if !tv.Truth() {
		t.Fatal("Expected Truth = 'true'; got 'false'")
	}

	tv = getValueType(type_, attrname, "")

	if tv.Truth() {
		t.Fatal("Expected Truth = 'false'; got 'true'")
	}

	tv = getValueType(type_, "", "")

	if tv.Truth() {
		t.Fatal("Expected Truth = 'false'; got 'true'")
	}
}

func TestSetFieldInvalidType(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	err := tv.SetField(attrname, starlark.MakeInt(0))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestSetFieldOnFrozenObj(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	tv.Freeze()
	err := tv.SetField(attrname, starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestSetKeyOnFrozenObj(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	tv.Freeze()
	err := tv.SetKey(starlark.String(attrname), starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestSetFieldOnInvalidAttName(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	err := tv.SetField("this attr name does not exist", starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestAttrNames(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	attrvalue := "value"
	tvs := []*TestValue{
		getValueType(type_, attrname, attrvalue),
		getMapValueType(type_, attrname, attrvalue),
	}
	for _, tv := range tvs {
		attrnames := tv.AttrNames()
		if len(attrnames) != 1 {
			t.Fatalf("Expected 1 attr; found: %d in '%v'", len(attrnames), attrnames)
		}
		if attrnames[0] != attrname {
			t.Fatalf("Expected attr with name '%s'; found: %s", attrname, attrnames[0])
		}
		value, err := tv.Attr(attrname)
		if err != nil {
			t.Fatalf("Unexpected error '%v'", err)
		}
		str, ok := starlark.AsString(value)
		if !ok {
			t.Fatalf("Error in string conversion for '%T'", value)
		}
		if str != attrvalue {
			t.Fatalf("Expected '%s'; got '%s'", attrvalue, str)
		}
		expectstr := "TestValue{attrname=\"value\"}"
		if tv.String() != expectstr {
			t.Fatalf("expected '%s'; got '%s'", expectstr, tv.String())
		}
	}
}

func TestHashError(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	_, err := tv.Hash()
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestGetByInvalidType(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	_, _, err := tv.Get(starlark.MakeInt(1))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestGetMissingAttr(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	_, found, err := tv.Get(starlark.String("this attr does not exist"))
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if found {
		t.Fatal("Attr was unexpectedly found")
	}
}

func TestSetKeyWithInvalidType(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "")
	err := tv.SetKey(starlark.MakeInt(1), starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; none found")
	}

	err = tv.SetKey(starlark.String(attrname), starlark.MakeInt(1))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestSetKey(t *testing.T) {
	type_ := "TestValue"
	attrname := "attrname"
	tv := getValueType(type_, attrname, "something")

	err := tv.SetKey(starlark.String(attrname), starlark.String(""))
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}

	value, err := tv.Attr(attrname)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	strvalue, ok := starlark.AsString(value)
	if !ok {
		t.Fatalf("Expected string; found '%T'", value)
	}
	if strvalue != "" {
		t.Fatalf("Expected '%s'; got '%s'", "", strvalue)
	}
}

func TestIntValidator(t *testing.T) {
	type_, attrname := "TestValue", "attrname"
	tv := &TestValue{
		NewEKValue(
			type_,
			[]string{attrname},
			map[string]starlark.Value{},
			map[string]PreProcessFn{
				attrname: AssertInt,
			},
			NoMaskFn,
		),
	}
	err := tv.SetField(attrname, starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; none found")
	}

	err = tv.SetField(attrname, starlark.MakeInt(1))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
