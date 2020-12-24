// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"
	"time"

	"go.starlark.net/starlark"
)

func getTransactions(t *testing.T, fpath string) ([]*Transaction, error) {
	sc, err := New(fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	cred := NewCredentials("mr_user", "a_password", "an_account")
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

func Test_004_Errors(t *testing.T) {
	root := "../../test/ekm/vdefault/004_transactions/"
	testErrorFiles(t, root, func(path string) error {
		_, err := getTransactions(t, path)
		return err
	})
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

func Test_004_Keyword_Params(t *testing.T) {
	fpath := "../../test/ekm/vdefault/004_transactions/OK_keyword_params.ekm"
	txs, err := getTransactions(t, fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(txs) != 1 {
		t.Fatalf("Expected transactions slice with length 1; found '%v'", txs)
	}
	expect := starlark.MakeInt(100)
	got, err := txs[0].Attr(txAmount)
	equal, err := starlark.Equal(got, expect)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !equal {
		t.Fatalf("Expected '%v'; got '%v'", expect, got)
	}
}
