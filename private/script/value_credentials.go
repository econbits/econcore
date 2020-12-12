// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

type Credentials struct {
	username starlark.String
	pwd      starlark.String
	account  starlark.String
}

// New function

func NewCredentials(username string, pwd string, account string) Credentials {
	return Credentials{
		starlark.String(username),
		starlark.String(pwd),
		starlark.String(account),
	}
}

// Implementing starlark Value interface

func (cred Credentials) String() string {
	return cred.username.String() + " credentials"
}

func (cred Credentials) Type() string {
	return "Credentials"
}

func (cred Credentials) Freeze() {
	// Credentials are already immutable
}

func (cred Credentials) Truth() starlark.Bool {
	return len(cred.username) > 0 || len(cred.pwd) > 0 || len(cred.account) > 0
}

func (cred Credentials) Hash() (uint32, error) {
	return 0, fmt.Errorf("Credentials can't be hashed")
}

// Implements starlark HasAttrs interface

func (cred Credentials) Attr(name string) (starlark.Value, error) {
	if name == "username" {
		return cred.username, nil
	}
	if name == "pwd" {
		return cred.pwd, nil
	}
	if name == "account" {
		return cred.account, nil
	}
	return nil, nil
}

func (cred Credentials) AttrNames() []string {
	return []string{"username", "pwd", "account"}
}
