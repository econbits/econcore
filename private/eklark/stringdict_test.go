// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestGetStringFromDict(t *testing.T) {
	key := "key"
	expect := "test"
	sd := StringDict{key: starlark.String(expect)}
	got, err := sd.GetString(key)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != expect {
		t.Fatalf("expected %s; got %s", expect, got)
	}
}

func TestGetMissingStringFromDict(t *testing.T) {
	key := "key"
	sd := StringDict{}
	_, err := sd.GetString(key)
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetNoneAsStringFromDict(t *testing.T) {
	key := "key"
	sd := StringDict{key: starlark.None}
	_, err := sd.GetString(key)
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetIntAsStringFromDict(t *testing.T) {
	key := "key"
	sd := StringDict{key: starlark.MakeInt(1)}
	_, err := sd.GetString(key)
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetNonOptionalStringFromDict(t *testing.T) {
	key := "key"
	expect := "test"
	sd := StringDict{key: starlark.String(expect)}
	got, err := sd.GetStringOr(key, "alt")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != expect {
		t.Fatalf("expected %s; got %s", expect, got)
	}
}

func TestGetMissingOptionalStringFromDict(t *testing.T) {
	key := "key"
	sd := StringDict{}
	_, err := sd.GetStringOr(key, "alt")
	if err == nil {
		t.Fatal("expected error; got none")
	}
}

func TestGetOptionalStringFromDict(t *testing.T) {
	key := "key"
	expect := "alt"
	sd := StringDict{key: starlark.None}
	got, err := sd.GetStringOr(key, expect)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != expect {
		t.Fatalf("expected %s; got %s", expect, got)
	}
}

func TestGetIntAsOptionalStringFromDict(t *testing.T) {
	key := "key"
	sd := StringDict{key: starlark.MakeInt(1)}
	_, err := sd.GetStringOr(key, "alt")
	if err == nil {
		t.Fatal("expected error; got none")
	}
}
