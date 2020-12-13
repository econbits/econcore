// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func Test_002_Login_Empty(t *testing.T) {
	fpath := "../../test/ekm/vdefault/002_login/empty.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	cred := Credentials{username: "mr_user"}
	_, err = sc.Login(cred)
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}

func Test_002_Login_fail(t *testing.T) {
	fpath := "../../test/ekm/vdefault/002_login/fail.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	cred := Credentials{username: "mr_user"}
	_, err = sc.Login(cred)
	if err == nil {
		t.Errorf("Expected error; none found")
		return
	}
	if !strings.Contains(err.Error(), "[mr_user]") {
		t.Errorf("Expected '[mr_user]' in error; found: %s", err.Error())
	}
}

func Test_002_Login_No_Params(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/002_login/no_params.ekm",
		"../../test/ekm/vdefault/002_login/too_many_params.ekm",
	}
	for _, fpath := range fpaths {
		sc, err := New(fpath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		cred := Credentials{username: "mr_user"}
		_, err = sc.Login(cred)
		if err == nil {
			t.Errorf("Expected error; none found")
			return
		}
		if !strings.Contains(err.Error(), "function login") {
			t.Errorf("Expected 'function login' in error; found: %s", err.Error())
		}
	}
}

func Test_002_Login_success(t *testing.T) {
	fpath := "../../test/ekm/vdefault/002_login/success.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	cred := Credentials{username: "mr_user"}
	session, err := sc.Login(cred)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if session.Truth() {
		t.Errorf("Expected empty session; found: %v", session)
	}
}

func Test_002_Login_Wrong_Returns(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/002_login/wrong_return_type.ekm",
		"../../test/ekm/vdefault/002_login/no_session.ekm",
	}
	for _, fpath := range fpaths {
		sc, err := New(fpath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		cred := Credentials{username: "mr_user"}
		_, err = sc.Login(cred)
		if err == nil {
			t.Errorf("Expected error; none found")
			return
		}
		if !strings.Contains(err.Error(), "instead of a session") {
			t.Errorf("Expected 'instead of a session' in error; found: %s", err.Error())
		}
	}
}

func Test_002_Login_Set_Session_Param(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/002_login/set_session_param_1.ekm",
		"../../test/ekm/vdefault/002_login/set_session_param_2.ekm",
	}
	for _, fpath := range fpaths {
		sc, err := New(fpath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		cred := Credentials{username: "mr_user"}
		session, err := sc.Login(cred)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
		v, found, err := session.Get(starlark.String("key"))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
		if !found {
			t.Errorf("Expected key 'key' in session")
		}
		vs, ok := starlark.AsString(v)
		if !ok {
			t.Errorf("Expected string type for value; found '%T'", v)
		}
		if vs != "value" {
			t.Errorf("Expected 'value'; found '%v'", vs)
		}
	}
}

func Test_002_Login_Wrong_Session_Use(t *testing.T) {
	fpaths := []string{
		"../../test/ekm/vdefault/002_login/set_wrong_session_key.ekm",
		"../../test/ekm/vdefault/002_login/session_access_error_1.ekm",
		"../../test/ekm/vdefault/002_login/session_access_error_2.ekm",
		"../../test/ekm/vdefault/002_login/session_with_unnamed_param.ekm",
	}
	for _, fpath := range fpaths {
		sc, err := New(fpath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		cred := Credentials{username: "mr_user"}
		_, err = sc.Login(cred)
		if err == nil {
			t.Errorf("Expected error; none found")
		}
	}
}
