// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"math/big"
	"testing"
	"time"

	"github.com/econbits/econkit/private/ekres/credentials"
	"github.com/econbits/econkit/private/ekres/currency"
	"github.com/econbits/econkit/private/ekres/money"
	"github.com/econbits/econkit/private/ekres/transaction"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func getTransactions(t *testing.T, fpath string) ([]*transaction.Transaction, error) {
	sc, err := New(fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	cred := credentials.New("mr_user", "a_password", "an_account")
	session, err := sc.Login(cred)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	accounts, err := sc.Accounts(session)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(accounts) != 1 {
		t.Fatalf("Expected 1 account; found: %v", accounts)
	}
	transactions, err := sc.Transactions(session, accounts[0], time.Now())
	return transactions, err
}

func Test_004_Scripts(t *testing.T) {
	dpath := "../../test/ekm/vdefault/004_transactions/"
	epilogue := starlark.StringDict{}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		func(path string, epilogue starlark.StringDict, load testscript.LoadFn) error {
			_, err := getTransactions(t, path)
			return err
		},
		testscript.Fail,
	)
}

func Test_004_Empty_List(t *testing.T) {
	fpath := "../../test/ekm/vdefault/004_transactions/OK_empty_list.ekm"
	txs, err := getTransactions(t, fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(txs) > 0 {
		t.Fatalf("Expected empty transactions slice; found '%v'", txs)
	}
}

func Test_004_OK_Transactions(t *testing.T) {
	fpath := "../../test/ekm/vdefault/004_transactions/OK_transactions.ekm"
	txs, err := getTransactions(t, fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(txs) != 1 {
		t.Fatalf("Expected transactions slice with length 1; found '%v'", txs)
	}
	expect := money.New(big.NewInt(1), currency.MustGet("EUR"))
	got := txs[0].Value()
	if !expect.Equal(got) {
		t.Fatalf("Expected '%v'; got '%v'", expect, got)
	}
}
