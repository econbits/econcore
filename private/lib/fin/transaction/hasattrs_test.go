// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"math/big"
	"testing"
	"time"

	"github.com/econbits/econkit/private/lib/account/walletaccount"
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/lib/fin/money"
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

func TestHasAttrsMustGetTransaction(t *testing.T) {
	attrname := "attr"
	wallet := walletaccount.New("id", "name", "provider")
	attrvalue := New(
		wallet,
		wallet,
		money.New(big.NewInt(1), currency.MustGet("EUR")),
		datetime.NewFromTime(time.Now()),
		nil,
		"purpose",
	)

	tv := getTestHasAttrsValue(attrname, attrvalue)
	gotvalue := HasAttrsMustGetTransaction(tv, attrname)
	if !gotvalue.Equal(attrvalue) {
		t.Fatalf("Expected %v; found '%v'", attrvalue, gotvalue)
	}
}

func TestHasAttrsMustGetTransactionMissingAttr(t *testing.T) {
	attrname := "attr"
	wallet := walletaccount.New("id", "name", "provider")
	attrvalue := New(
		wallet,
		wallet,
		money.New(big.NewInt(1), currency.MustGet("EUR")),
		datetime.NewFromTime(time.Now()),
		nil,
		"purpose",
	)

	tv := getTestHasAttrsValue(attrname, attrvalue)

	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	HasAttrsMustGetTransaction(tv, "this attr does not exist")
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

	HasAttrsMustGetTransaction(tv, attrname)
}
