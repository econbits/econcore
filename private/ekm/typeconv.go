// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"fmt"

	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/transaction"
	"go.starlark.net/starlark"
)

// Integer Range to Value Range
func IRtoVR(ilist []int) []starlark.Value {
	li := len(ilist)
	values := make([]starlark.Value, li)
	for i := 0; i < li; i++ {
		values[i] = starlark.MakeInt(ilist[i])
	}
	return values
}

// Value Range to Integer Range
func VRtoIR(vlist []starlark.Value) ([]int, error) {
	lv := len(vlist)
	ilist := make([]int, lv)
	for i := 0; i < lv; i++ {
		ivalue, err := starlark.AsInt32(vlist[i])
		ilist[i] = ivalue
		if err != nil {
			return nil, err
		}
	}
	return ilist, nil
}

// List of accounts to Account Range
func LtoAR(list *starlark.List) ([]*account.Account, error) {
	llist := list.Len()
	alist := make([]*account.Account, llist)
	for i := 0; i < llist; i++ {
		v := list.Index(i)
		va, ok := v.(*account.Account)
		if !ok {
			return nil, fmt.Errorf("Expected List of Accounts; found: %T", v)
		}
		alist[i] = va
	}
	return alist, nil
}

// List of transactions to Transaction Range
func LtoTR(list *starlark.List) ([]*transaction.Transaction, error) {
	llist := list.Len()
	tlist := make([]*transaction.Transaction, llist)
	for i := 0; i < llist; i++ {
		v := list.Index(i)
		vt, ok := v.(*transaction.Transaction)
		if !ok {
			return nil, fmt.Errorf("Expected List of Transactions; found: %T", v)
		}
		tlist[i] = vt
	}
	return tlist, nil
}
