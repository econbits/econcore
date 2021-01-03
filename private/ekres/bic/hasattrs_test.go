// Copyright (C) 2021  Germán Fuentes Capella

package bic

import (
	"testing"

	"github.com/econbits/econkit/private/eklark"
	"go.starlark.net/starlark"
)

type TestValue struct {
	eklark.EKValue
}

func getTestHasAttrsValue(attrname string, attrvalue starlark.Value) starlark.HasAttrs {
	type_ := "TestValue"
	tv := &TestValue{
		eklark.NewEKValue(
			type_,
			[]string{attrname},
			map[string]starlark.Value{
				attrname: attrvalue,
			},
			map[string]eklark.ValidateFn{},
			map[string]eklark.FormatterFn{},
			eklark.NoMaskFn,
		),
	}
	return tv
}

func TestHasAttrsMustGetBIC(t *testing.T) {
	attrname := "attr"
	attrvalue, err := Parse("DEUTDEFFXXX")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	tv := getTestHasAttrsValue(attrname, attrvalue)
	gotvalue := HasAttrsMustGetBIC(tv, attrname)
	if !gotvalue.Equal(attrvalue) {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetBICMissingAttr(t *testing.T) {
	attrname := "attr"
	attrvalue, err := Parse("DEUTDEFFXXX")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetBIC(tv, "this attr does not exist")
}

func TestHasAttrsMustGetBICNotABIC(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetBIC(tv, attrname)
}
