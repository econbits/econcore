// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestGetStringFromDict(t *testing.T) {
	key := "key"
	expect := "test"
	sd := starlark.StringDict{key: starlark.String(expect)}
	got, err := StringDictGetString(sd, key)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != expect {
		t.Fatalf("expected %s; got %s", expect, got)
	}
}

func TestGetMissingStringFromDict(t *testing.T) {
	key := "key"
	sd := starlark.StringDict{}
	_, err := StringDictGetString(sd, key)
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetNoneAsStringFromDict(t *testing.T) {
	key := "key"
	sd := starlark.StringDict{key: starlark.None}
	_, err := StringDictGetString(sd, key)
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetIntAsStringFromDict(t *testing.T) {
	key := "key"
	sd := starlark.StringDict{key: starlark.MakeInt(1)}
	_, err := StringDictGetString(sd, key)
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetNonOptionalStringFromDict(t *testing.T) {
	key := "key"
	expect := "test"
	sd := starlark.StringDict{key: starlark.String(expect)}
	got, err := StringDictGetStringOr(sd, key, "alt")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != expect {
		t.Fatalf("expected %s; got %s", expect, got)
	}
}

func TestGetMissingOptionalStringFromDict(t *testing.T) {
	key := "key"
	sd := starlark.StringDict{}
	_, err := StringDictGetStringOr(sd, key, "alt")
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetOptionalStringFromDict(t *testing.T) {
	key := "key"
	expect := "alt"
	sd := starlark.StringDict{key: starlark.None}
	got, err := StringDictGetStringOr(sd, key, expect)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != expect {
		t.Fatalf("expected %s; got %s", expect, got)
	}
}

func TestGetIntAsOptionalStringFromDict(t *testing.T) {
	key := "key"
	sd := starlark.StringDict{key: starlark.MakeInt(1)}
	_, err := StringDictGetStringOr(sd, key, "alt")
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetBool(t *testing.T) {
	key1 := "key1"
	key2 := "key2"
	sd := starlark.StringDict{key1: starlark.True, key2: starlark.String("a")}
	for _, k := range []string{key1, key2} {
		ok, err := StringDictGetBool(sd, k)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if !ok {
			t.Fatal("expected true; got false")
		}
	}
}

func TestGetBoolFromNone(t *testing.T) {
	key := "key1"
	sd := starlark.StringDict{key: starlark.None}
	value, err := StringDictGetBool(sd, key)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if value {
		t.Fatal("expected false; found true")
	}
}

func TestGetBoolFromNil(t *testing.T) {
	key := "key1"
	sd := starlark.StringDict{key: nil}
	_, err := StringDictGetBool(sd, key)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}

func TestGetMissingBool(t *testing.T) {
	k := "key1"
	sd := starlark.StringDict{}
	_, err := StringDictGetBool(sd, k)
	if err == nil {
		t.Fatal("expected error; none found")
	}
}
