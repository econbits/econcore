// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"reflect"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestNewCredentials(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := NewCredentials(username, pwd, account)
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

func TestWrongAttrInCredentials(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := NewCredentials(username, pwd, account)
	v, err := c.Attr("this_attr_does_not_exist")
	if err != nil {
		t.Errorf("Unexpected error; got '%v'", err)
	}
	if v != nil {
		t.Errorf("Unexpected value; got '%v'", v)
	}
}

func TestCredentialsHashError(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := NewCredentials(username, pwd, account)
	_, err := c.Hash()
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}

func TestCredentialsTruth(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := NewCredentials(username, pwd, account)
	if !c.Truth() {
		t.Errorf("Expected Credentials Truth=true; got false")
	}
	c = Credentials{}
	if c.Truth() {
		t.Errorf("Expected Credentials Truth=false; got true")
	}
}

func TestCredentialsEasyMethods(t *testing.T) {
	username, pwd, account := "username", "pwd", "account"
	c := NewCredentials(username, pwd, account)

	// this should do nothing. As long as it does not panic, we are fine
	c.Freeze()

	if c.Type() != "Credentials" {
		t.Errorf("Expected Credentials Type='Credentials'; got %s", c.Type())
	}

	if !strings.Contains(c.String(), "username") {
		t.Errorf("Expected Credentials String to contain the username; got %s", c.String())
	}
}
