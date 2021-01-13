// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"reflect"
	"testing"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func Test_001_Errors(t *testing.T) {
	dpath := "../../test/ekm/vdefault/001_globals/"
	epilogue := starlark.StringDict{}
	testscript.TestingRun(
		t,
		dpath,
		epilogue,
		testscript.LoadEmptyFn,
		func(path string, epilogue starlark.StringDict, load testscript.LoadFn) error {
			_, err := New(path)
			return err
		},
		testscript.Fail,
	)
}

func Test_001_Meta_Full(t *testing.T) {
	fpath := "../../test/ekm/vdefault/001_globals/OK_full.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if sc.Description() != "this is a test module" {
		t.Fatalf("Expected description: 'this is a test module'; got '%s'", sc.Description())
	}
	if sc.URL() != "https://econkit.org" {
		t.Fatalf("Expected URL: 'https://econkit.org'; got '%s'", sc.URL())
	}
	if sc.License() != "MIT" {
		t.Fatalf("Expected License: 'MIT'; got '%s'", sc.License())
	}
	authors := []string{"Mr. T", "Mr. L"}
	if !reflect.DeepEqual(sc.Authors(), authors) {
		t.Fatalf("Expected Authors: '%v'; got '%v'", authors, sc.Authors())
	}
}

func Test_001_Meta_Empty(t *testing.T) {
	fpath := "../../test/ekm/vdefault/001_globals/OK_empty.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if sc.Description() != defDescription {
		t.Fatalf("Expected description: '%s'; got '%s'", defDescription, sc.Description())
	}
	if sc.URL() != defUrl {
		t.Fatalf("Expected URL: '%s'; got '%s'", defUrl, sc.URL())
	}
	if sc.License() != defLicense {
		t.Fatalf("Expected License: '%s'; got '%s'", defLicense, sc.License())
	}
	if !reflect.DeepEqual(sc.Authors(), defAuthors) {
		t.Fatalf("Expected Authors: '%v'; got '%v'", defAuthors, sc.Authors())
	}
}

func Test_001_Meta_Corrupt_Values(t *testing.T) {
	fpath := "../../test/ekm/vdefault/001_globals/OK_full.ekm"
	sc, err := New(fpath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if sc.Description() != "this is a test module" {
		t.Fatalf("Expected description: 'this is a test module'; got '%s'", sc.Description())
	}
	sc.globals[hDescription] = starlark.MakeInt(1)
	if sc.Description() != defDescription {
		t.Fatalf("Expected description: '%s'; got '%s'", defDescription, sc.Description())
	}

	authors := []string{"Mr. T", "Mr. L"}
	if !reflect.DeepEqual(sc.Authors(), authors) {
		t.Fatalf("Expected Authors: '%v'; got '%v'", authors, sc.Authors())
	}
	sc.globals[hAuthors] = starlark.MakeInt(1)
	if !reflect.DeepEqual(sc.Authors(), defAuthors) {
		t.Fatalf("Expected Authors: '%v'; got '%v'", defAuthors, sc.Authors())
	}
}
