// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"math/big"
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertString(t *testing.T) {
	var value starlark.Value
	value = starlark.String("")
	err := AssertString(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	value = starlark.MakeInt(1)
	err = AssertString(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestAssertInt(t *testing.T) {
	var value starlark.Value
	value = starlark.MakeInt(1)
	err := AssertInt(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	value = starlark.String("")
	err = AssertInt(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestAssertInt32(t *testing.T) {
	var value starlark.Value
	value = starlark.MakeInt(1)
	err := AssertInt32(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	value = starlark.MakeInt64(int64(^uint64(0) >> 1))
	err = AssertInt32(value)
	if err == nil {
		t.Fatal("[Int64 -> Int] expected error; none found")
	}

	value = starlark.String("")
	err = AssertInt32(value)
	if err == nil {
		t.Fatal("[string -> int] expected error; none found")
	}
}

func TestAssertInt64(t *testing.T) {
	var value starlark.Value
	value = starlark.MakeInt(1)
	err := AssertInt64(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	bint := big.NewInt(int64(^uint64(0) >> 1))
	zint := big.NewInt(0)
	value = starlark.MakeBigInt(zint.Add(bint, bint))
	err = AssertInt64(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}

	value = starlark.String("")
	err = AssertInt64(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestAssertUint64(t *testing.T) {
	var value starlark.Value
	value = starlark.MakeUint64(1)
	err := AssertUint64(value)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	//bint := big.NewInt(int64(^uint64(0) >> 1))
	zint := big.NewInt(-1)
	value = starlark.MakeBigInt(zint)
	err = AssertUint64(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}

	value = starlark.String("")
	err = AssertUint64(value)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
