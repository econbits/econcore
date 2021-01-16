// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"testing"

	"github.com/econbits/econkit/private/lib/auth/credentials"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func Test_002_Scripts(t *testing.T) {
	dpath := "../../test/ekm/vdefault/002_login/"
	epilogue := starlark.StringDict{}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		func(path string, epilogue starlark.StringDict, load testscript.LoadFn) error {
			sc, err := New(path)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			cred := credentials.New("mr_user", "a_password", "an_account")
			_, err = sc.Login(cred)
			return err
		},
		testscript.Fail,
	)
}

func Test_002_Login_success(t *testing.T) {
	fpath := "../../test/ekm/vdefault/002_login/OK_success.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	cred := credentials.New("mr_user", "a_password", "an_account")
	session, err := sc.Login(cred)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if session.Truth() {
		t.Fatalf("Expected empty session; found: %v", session)
	}
}

func Test_002_Login_Set_Session_Param(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/002_login/OK_set_session_param_1.ekm",
		"../../test/ekm/vdefault/002_login/OK_set_session_param_2.ekm",
	}
	for _, fpath := range fpaths {
		sc, err := New(fpath)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		cred := credentials.New("mr_user", "a_password", "an_account")
		session, err := sc.Login(cred)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		v, found, err := session.Get(starlark.String("key"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !found {
			t.Fatalf("Expected key 'key' in session")
		}
		vs, ok := starlark.AsString(v)
		if !ok {
			t.Fatalf("Expected string type for value; found '%T'", v)
		}
		if vs != "value" {
			t.Fatalf("Expected 'value'; found '%v'", vs)
		}
	}
}
