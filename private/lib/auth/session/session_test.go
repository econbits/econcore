// Copyright (C) 2020  Germán Fuentes Capella

package session

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestSessionFreeze(t *testing.T) {
	s := New()
	err := s.SetKey(starlark.String("key"), starlark.String("value"))
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	s.Freeze()
	err = s.SetKey(starlark.String("key"), starlark.String("value"))
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}

func TestSessionAttr(t *testing.T) {
	s := New()
	key := starlark.String("key")
	expect := starlark.String("value")
	err := s.SetKey(key, expect)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	got, found, err := s.Get(key)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if !found {
		t.Fatalf("Key '%v' was not found in session '%v'", key, s)
	}
	equal, err := starlark.Equal(got, expect)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if !equal {
		t.Fatalf("Expected '%v'; got '%v'", expect, got)
	}
	//
	got, err = s.Attr("key")
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	equal, err = starlark.Equal(got, expect)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if !equal {
		t.Fatalf("Expected '%v'; got '%v'", expect, got)
	}
	//
	attrs := s.AttrNames()
	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute; got '%v' with len %d", attrs, len(attrs))
	}
	if attrs[0] != "key" {
		t.Fatalf("Expected attribute 'key'; got '%v'", attrs[0])
	}
}

func TestSessionNonExistingAttr(t *testing.T) {
	s := New()
	v, err := s.Attr("this_attr_does_not_exist")
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	if v != nil {
		t.Fatalf("Unexpected value '%v'", v)
	}
}

func TestSessionEmpty(t *testing.T) {
	s := New()
	key := starlark.String("key")
	value := starlark.String("value")
	err := s.SetKey(key, value)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
	expect := "Session{key=\"value\"}"
	got := s.String()
	if got != expect {
		t.Fatalf("Expected session '%v'; got '%s'", expect, got)
	}
}
