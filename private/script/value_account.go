//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

type Account struct {
	name starlark.String
	typ  starlark.String
	iban starlark.String
	bic  starlark.String
}

// Implementing starlark Value interface

func (a Account) String() string {
	return fmt.Sprintf("Account{name=%v, typ=%v, iban=%v, bic=%v}", a.name, a.typ, a.iban, a.bic)
}

func (a Account) Type() string {
	return "Account"
}

func (a Account) Freeze() {
}

func (a Account) Truth() starlark.Bool {
	return a.name.Len() > 0 || a.typ.Len() > 0 || a.iban.Len() > 0 || a.bic.Len() > 0
}

func (a Account) Hash() (uint32, error) {
	return 0, fmt.Errorf("Account can't be hashed")
}

// Implements starlark HasAttrs interface

func (a Account) Attr(name string) (starlark.Value, error) {
	if name == "name" {
		return a.name, nil
	}
	if name == "typ" {
		return a.typ, nil
	}
	if name == "iban" {
		return a.iban, nil
	}
	if name == "bic" {
		return a.bic, nil
	}
	return nil, nil
}

func (a Account) AttrNames() []string {
	return []string{"name", "typ", "iban", "bic"}
}
