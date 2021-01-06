// Copyright (C) 2021  Germán Fuentes Capella

package eklark

import (
	"fmt"
	"reflect"
	"testing"

	"go.starlark.net/starlark"
)

func getValueTypeWithAlias(t *testing.T, type_, attrname, attrvalue, alias string) *TestValue {
	tv := &TestValue{
		NewEKValue(
			type_,
			[]string{attrname},
			map[string]starlark.Value{
				attrname: starlark.String(attrvalue),
			},
			map[string]PreProcessFn{
				attrname: AssertString,
			},
			NoMaskFn,
		),
	}
	err := tv.SetAlias(alias, attrname)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	return tv
}
func TestValueSetAliasError(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	err := v.SetAlias(alias, "this alias does not exist")
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}

func TestValueAliasString(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	// String
	expect := fmt.Sprintf("%s{%s=\"%s\", %s=\"%s\"}", testType, attr, value, alias, value)
	got := v.String()
	if expect != got {
		t.Fatalf("expected '%s'; found '%s'", expect, got)
	}
}

func TestValueAliasAttr(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	vexpect, err := v.Attr(attr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	vgot, err := v.Attr(alias)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if vgot != vexpect {
		t.Fatalf("expected %v; got %v", vexpect, vgot)
	}
}

func TestValueAliasAttrNames(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	expectAttrs := []string{attr, alias}
	gotAttrs := v.AttrNames()
	if !reflect.DeepEqual(expectAttrs, gotAttrs) {
		t.Fatalf("expected %v; got %v", expectAttrs, gotAttrs)
	}
}

func TestValueAliasSetField(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	newvalue := "new value"
	vnew := starlark.String(newvalue)
	err := v.SetField(attr, vnew)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	vgot, err := v.Attr(alias)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if vgot != vnew {
		t.Fatalf("expected %v; got %v", vnew, vgot)
	}
	vold := starlark.String(value)
	err = v.SetField(alias, vold)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	vgot, err = v.Attr(attr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if vgot != vold {
		t.Fatalf("expected %v; got %v", vold, vgot)
	}
}

func TestValueAliasGet(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	vexpect := starlark.String(value)
	vgot, found, err := v.Get(starlark.String(alias))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if !found {
		t.Fatal("expected found")
	}
	if vgot != vexpect {
		t.Fatalf("expected %v; got %v", vexpect, vgot)
	}
}

func TestValueAliasSetKey(t *testing.T) {
	testType, attr, value, alias := "TestAlias", "attr", "value", "alias"
	v := getValueTypeWithAlias(t, testType, attr, value, alias)
	vold := starlark.String(value)
	newvalue := "new value"
	vnew := starlark.String(newvalue)
	err := v.SetKey(starlark.String(attr), vnew)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	vgot, err := v.Attr(alias)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if vgot != vnew {
		t.Fatalf("expected %v; got %v", vnew, vgot)
	}
	err = v.SetKey(starlark.String(alias), vold)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	vgot, err = v.Attr(attr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if vgot != vold {
		t.Fatalf("expected %v; got %v", vold, vgot)
	}
}
