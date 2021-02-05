// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"github.com/econbits/econkit/private/lib/account/account"
	"github.com/econbits/econkit/private/lib/fin/transaction"
	"go.starlark.net/starlark"
)

// List of accounts to Account Range
func LtoAR(list *starlark.List) ([]*account.Account, bool) {
	llist := list.Len()
	alist := make([]*account.Account, llist)
	for i := 0; i < llist; i++ {
		v := list.Index(i)
		va, ok := v.(*account.Account)
		if !ok {
			return nil, false
		}
		alist[i] = va
	}
	return alist, true
}

// List of transactions to Transaction Range
func LtoTR(list *starlark.List) ([]*transaction.Transaction, bool) {
	llist := list.Len()
	tlist := make([]*transaction.Transaction, llist)
	for i := 0; i < llist; i++ {
		v := list.Index(i)
		vt, ok := v.(*transaction.Transaction)
		if !ok {
			return nil, false
		}
		tlist[i] = vt
	}
	return tlist, true
}
