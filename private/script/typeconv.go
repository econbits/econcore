//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

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

// Value to Integer Range
func VtoIR(value starlark.Value) ([]int, error) {
	vlist, ok := value.(*starlark.List)
	if !ok {
		return nil, fmt.Errorf("Expected *starlark.List; got %T", value)
	}
	return LtoIR(vlist)
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

// List of Integers to Integer Range
func LtoIR(list *starlark.List) ([]int, error) {
	llist := list.Len()
	nlist := make([]int, llist)
	for i := 0; i < llist; i++ {
		v := list.Index(i)
		vi, err := starlark.AsInt32(v)
		if err != nil {
			return nil, err
		}
		nlist[i] = vi
	}
	return nlist, nil
}
