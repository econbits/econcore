// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"

	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/credentials"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func getAccounts(t *testing.T, fpath string) ([]*account.Account, error) {
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
	return accounts, err
}

func Test_003_Scripts(t *testing.T) {
	dpath := "../../test/ekm/vdefault/003_accounts/"
	epilogue := starlark.StringDict{}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		func(path string, epilogue starlark.StringDict) error {
			_, err := getAccounts(t, path)
			return err
		},
		testscript.Fail,
	)
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
