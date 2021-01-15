// Copyright (C) 2021  Germán Fuentes Capella

package transaction

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/money"
	dtlib "github.com/econbits/econkit/private/lib/datetime"
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/lib/iso/currency"
	"github.com/econbits/econkit/private/slang"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/transaction/"
	fns := []*slang.Fn{
		TransactionFn,
		account.WalletFn,
		currency.CurrencyFn,
		money.MoneyFn,
	}
	epilogue := starlark.StringDict{}
	for _, fn := range fns {
		epilogue[fn.Name] = fn.Builtin()
	}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			if module == dtlib.Lib.Name {
				return dtlib.Lib.Load(), nil
			}
			return nil, fmt.Errorf("unknown module: %s", module)
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}

func TestTransactionAttributes(t *testing.T) {
	acc := account.NewWalletAccount("id", "name", "provider")
	value := money.New(big.NewInt(100), currency.MustGet("EUR"))
	dt := datetime.NewFromTime(time.Now())
	purpose := "purpose"
	tx := New(acc, acc, value, dt, dt, starlark.String(purpose))
	if !tx.Sender().Equal(acc) {
		t.Fatalf("expected %v; got %v", acc, tx.Sender())
	}
	if !tx.Receiver().Equal(acc) {
		t.Fatalf("expected %v; got %v", acc, tx.Receiver())
	}
	if !tx.Value().Equal(value) {
		t.Fatalf("expected %v; got %v", value, tx.Value())
	}
	if !tx.ValueDate().Equal(dt) {
		t.Fatalf("expected %v; got %v", dt, tx.ValueDate())
	}
	if !tx.BookingDate().Equal(dt) {
		t.Fatalf("expected %v; got %v", dt, tx.BookingDate())
	}
	if tx.Purpose() != purpose {
		t.Fatalf("expected %v; got %v", purpose, tx.Purpose())
	}
}
