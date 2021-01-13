// Copyright (C) 2021  Germán Fuentes Capella

package datetime

import (
	"testing"
	"time"

	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type TestValue struct {
	slang.EKValue
}

func getTestHasAttrsValue(attrname string, attrvalue starlark.Value) starlark.HasAttrs {
	type_ := "TestValue"
	tv := &TestValue{
		slang.NewEKValue(
			type_,
			[]string{attrname},
			map[string]starlark.Value{
				attrname: attrvalue,
			},
			map[string]slang.PreProcessFn{},
			slang.NoMaskFn,
		),
	}
	return tv
}

func TestHasAttrsMustGetDateTime(t *testing.T) {
	attrname := "attr"
	attrvalue := NewFromTime(time.Now())

	tv := getTestHasAttrsValue(attrname, attrvalue)
	gotvalue := HasAttrsMustGetDateTime(tv, attrname)
	if !gotvalue.Equal(attrvalue) {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetDateTimeMissingAttr(t *testing.T) {
	attrname := "attr"
	attrvalue := NewFromTime(time.Now())

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetDateTime(tv, "this attr does not exist")
}

func TestHasAttrsMustGetDateTimeNotADateTime(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetDateTime(tv, attrname)
}
