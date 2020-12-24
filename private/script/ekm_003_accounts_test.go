// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"

	"go.starlark.net/starlark"
)

func getAccounts(t *testing.T, fpath string) ([]*Account, error) {
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
	return accounts, err
}

func Test_003_Errors(t *testing.T) {
	root := "../../test/ekm/vdefault/003_accounts/"
	testErrorFiles(t, root, func(path string) error {
		_, err := getAccounts(t, path)
		return err
	})
}

func Test_003_Accounts_Empty_List(t *testing.T) {
	fpath := "../../test/ekm/vdefault/003_accounts/OK_empty_list.ekm"
	accounts, err := getAccounts(t, fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(accounts) > 0 {
		t.Fatalf("Expected empty accounts slice; found '%v'", accounts)
	}
}

func Test_003_Accounts_Empty_Account(t *testing.T) {
	fpath := "../../test/ekm/vdefault/003_accounts/OK_empty_account.ekm"
	accounts, err := getAccounts(t, fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(accounts) != 1 {
		t.Fatalf("Expected 1 account; found '%v'", accounts)
	}
	if accounts[0].Truth() {
		t.Fatalf("Expected '%v'.Truth() false; found true", accounts[0])
	}
}

func Test_003_Accounts_Name_Param(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/003_accounts/OK_positional_param.ekm",
		"../../test/ekm/vdefault/003_accounts/OK_keyword_param.ekm",
	}
	for _, fpath := range fpaths {
		accounts, err := getAccounts(t, fpath)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if len(accounts) != 1 {
			t.Fatalf("Expected 1 account; found '%v'", accounts)
		}
		if !accounts[0].Truth() {
			t.Fatalf("Expected '%v'.Truth() true; found false", accounts[0])
		}
		namevalue, err := accounts[0].Attr("name")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		namestr, ok := starlark.AsString(namevalue)
		if !ok {
			t.Fatalf("Unexpected error getting string value from '%v'", namevalue)
		}
		if namestr != "test name" {
			t.Fatalf("Expected name 'test name'; found '%v'", namestr)
		}
		for _, attrname := range accounts[0].AttrNames() {
			if attrname == "name" {
				continue
			}
			value, err := accounts[0].Attr(attrname)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			valuestr, ok := starlark.AsString(value)
			if !ok {
				t.Fatalf("Unexpected error getting string value from '%v'", value)
			}
			if valuestr != "" {
				t.Fatalf("Expected ''; found '%v'", valuestr)
			}
		}
		_, err = accounts[0].Attr("this attr does not exist")
		if err == nil {
			t.Fatal("Expected error; found none")
		}
	}
}
