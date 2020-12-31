// Copyright (C) 2020  Germán Fuentes Capella

package credentials

import (
	"reflect"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestNew(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	attrs := []string{username, pwd, account}
	if !reflect.DeepEqual(c.AttrNames(), attrs) {
		t.Errorf("Expected: '%v'; got: '%v'", attrs, c.AttrNames())
	}

	for _, attr := range attrs {
		v, err := c.Attr(attr)
		if err != nil {
			t.Errorf("Unexpected error; got '%v'", err)
		}
		vs, ok := starlark.AsString(v)
		if !ok {
			t.Errorf("Unexpected error getting string from '%v'", v)
		}
		if vs != attr {
			t.Errorf("Expected: '%v'; got: '%v'", attr, vs)
		}
	}
}

func TestGetWrongAttrInCredentials(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	v, err := c.Attr("this_attr_does_not_exist")
	if err == nil {
		t.Fatal("Expected error; got none")
	}
	if v != nil {
		t.Errorf("Unexpected value; got '%v'", v)
	}
}

func TestSetWrongAttrInCredentials(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	err := c.SetField("this_attr_does_not_exist", starlark.String(""))
	if err == nil {
		t.Fatal("Expected error; got none")
	}
}

func TestSetWrongValueInCredentials(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	err := c.SetField("username", starlark.MakeInt(1))
	if err == nil {
		t.Fatal("Expected error; got none")
	}
}

func TestCredentialsHashError(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	_, err := c.Hash()
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}

func TestCredentialsTruth(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	if !c.Truth() {
		t.Errorf("Expected Credentials Truth=true; got false")
	}
	c = New("", "", "")
	if c.Truth() {
		t.Errorf("Expected Credentials Truth=false; got true")
	}
}

func TestCredentialsFreeze(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)
	expect := "a username"
	err := c.SetField(username, starlark.String(expect))
	if err != nil {
		t.Fatalf("username update failed for '%v', with error: %v", c, err)
	}
	got, err := c.Attr(username)
	if err != nil {
		t.Fatalf("getting username failed for '%v', with error: %v", c, err)
	}
	equal, err := starlark.Equal(got, starlark.String(expect))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !equal {
		t.Fatalf("Expected '%v', got: %v", expect, got)
	}
	c.Freeze()
	err = c.SetField(username, starlark.String(""))
	if err == nil {
		t.Fatal("username update was successful, expected failure")
	}
}

func TestCredentialsType(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)

	if c.Type() != "Credentials" {
		t.Errorf("Expected Credentials Type='Credentials'; got %s", c.Type())
	}
}

func TestCredentialsString(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := New(username, pwd, account)

	if !strings.Contains(c.String(), "pwd=\"*****\"") {
		t.Errorf("Password in Credentials's string is unmasked: %s", c.String())
	}

	if !strings.Contains(c.String(), "account=\"*****\"") {
		t.Errorf("Account in Credentials's string is unmasked: %s", c.String())
	}
}
