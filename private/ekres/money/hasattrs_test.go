// Copyright (C) 2021  Germán Fuentes Capella

package money

import (
	"math/big"
	"testing"

	"github.com/econbits/econkit/private/lib/iso/currency"
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

func TestHasAttrsMustGetMoney(t *testing.T) {
	attrname := "attr"
	attrvalue := New(big.NewInt(1), currency.MustGet("EUR"))

	tv := getTestHasAttrsValue(attrname, attrvalue)
	gotvalue := HasAttrsMustGetMoney(tv, attrname)
	if !gotvalue.Equal(attrvalue) {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetMoneyMissingAttr(t *testing.T) {
	attrname := "attr"
	attrvalue := New(big.NewInt(1), currency.MustGet("EUR"))

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetMoney(tv, "this attr does not exist")
}

func TestHasAttrsMustGetMoneyNotMoney(t *testing.T) {
	attrname := "attr"
	attrvalue := starlark.String("")

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetMoney(tv, attrname)
}
