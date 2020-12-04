//Copyright (C) 2020  Germán Fuentes Capella

package account

import (
	"fmt"
)

type Account struct {
	id       fmt.Stringer
	name     string
	typ      Type
	provider fmt.Stringer
}

func New(id fmt.Stringer, name string, typ Type, provider fmt.Stringer) Account {
	return Account{id: id, name: name, typ: typ, provider: provider}
}

func NewIBANAccount(name string, typ Type, iban IBAN, bic BIC) Account {
	return New(iban, name, typ, bic)
}

func (acc Account) Id() fmt.Stringer {
	return acc.id
}

func (acc Account) Name() string {
	return acc.name
}

func (acc Account) Type() Type {
	return acc.typ
}

func (acc Account) Provider() fmt.Stringer {
	return acc.provider
}
