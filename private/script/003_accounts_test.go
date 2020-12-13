//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"

	"go.starlark.net/starlark"
)

func getAccounts(t *testing.T, fpath string) ([]Account, error) {
	sc, err := New(fpath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	cred := Credentials{username: "mr_user"}
	session, err := sc.Login(cred)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	accounts, err := sc.Accounts(session)
	return accounts, err
}

func Test_003_Accounts_Errors(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/003_accounts/missing_function.ekm",
		"../../test/ekm/vdefault/003_accounts/no_session_param.ekm",
		"../../test/ekm/vdefault/003_accounts/too_many_params.ekm",
		"../../test/ekm/vdefault/003_accounts/return_none.ekm",
		"../../test/ekm/vdefault/003_accounts/int_list.ekm",
		"../../test/ekm/vdefault/003_accounts/positional_param_error.ekm",
		"../../test/ekm/vdefault/003_accounts/positional_and_keyword_param_conflict.ekm",
		"../../test/ekm/vdefault/003_accounts/keyword_param_error.ekm",
	}
	for _, fpath := range fpaths {
		_, err := getAccounts(t, fpath)
		if err == nil {
			t.Errorf("Expected error; none found")
		}
	}
}

func Test_003_Accounts_Empty_List(t *testing.T) {
	fpath := "../../test/ekm/vdefault/003_accounts/empty_list.ekm"
	accounts, err := getAccounts(t, fpath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(accounts) > 0 {
		t.Errorf("Expected empty accounts slice; found '%v'", accounts)
	}
}

func Test_003_Accounts_Empty_Account(t *testing.T) {
	fpath := "../../test/ekm/vdefault/003_accounts/empty_account.ekm"
	accounts, err := getAccounts(t, fpath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if len(accounts) != 1 {
		t.Errorf("Expected 1 account; found '%v'", accounts)
		return
	}
	if accounts[0].Truth() {
		t.Errorf("Expected '%v'.Truth() false; found true", accounts[0])
	}
}

func Test_003_Accounts_Name_Param(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/003_accounts/positional_param.ekm",
		"../../test/ekm/vdefault/003_accounts/keyword_param.ekm",
	}
	for _, fpath := range fpaths {
		accounts, err := getAccounts(t, fpath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
		if len(accounts) != 1 {
			t.Errorf("Expected 1 account; found '%v'", accounts)
			return
		}
		if !accounts[0].Truth() {
			t.Errorf("Expected '%v'.Truth() true; found false", accounts[0])
		}
		namevalue, err := accounts[0].Attr("name")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
		namestr, ok := starlark.AsString(namevalue)
		if !ok {
			t.Errorf("Unexpected error getting string value from '%v'", namevalue)
			return
		}
		if namestr != "test name" {
			t.Errorf("Expected name 'test name'; found '%v'", namestr)
		}
		for _, attrname := range accounts[0].AttrNames() {
			if attrname == "name" {
				continue
			}
			value, err := accounts[0].Attr(attrname)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			valuestr, ok := starlark.AsString(value)
			if !ok {
				t.Errorf("Unexpected error getting string value from '%v'", value)
			}
			if valuestr != "" {
				t.Errorf("Expected ''; found '%v'", valuestr)
			}
		}
		_, err = accounts[0].Attr("this attr does not exist")
		if err != nil {
			t.Errorf("Unxpected error; found: '%v'", err)
		}
	}
}
