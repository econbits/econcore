// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"reflect"
	"testing"

	"go.starlark.net/starlark"
)

func Test_001_Meta_Errors(t *testing.T) {
	fnames := []string{
		"../../test/ekm/vdefault/001_meta_authors_error_1.ekm",
		"../../test/ekm/vdefault/001_meta_authors_error_2.ekm",
		"../../test/ekm/vdefault/001_meta_description_error.ekm",
		"../../test/ekm/vdefault/001_meta_license_error.ekm",
		"../../test/ekm/vdefault/001_meta_url_error.ekm",
	}
	for _, fname := range fnames {
		_, err := New(fname)
		if err == nil {
			t.Errorf("[%s] Expected error; none found", fname)
		}
	}
}

func Test_001_Meta_Full(t *testing.T) {
	fname := "../../test/ekm/vdefault/001_meta_full.ekm"
	sc, err := New(fname)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sc.Description() != "this is a test module" {
		t.Errorf("Expected description: 'this is a test module'; got '%s'", sc.Description())
	}
	if sc.URL() != "https://econkit.org" {
		t.Errorf("Expected URL: 'https://econkit.org'; got '%s'", sc.URL())
	}
	if sc.License() != "MIT" {
		t.Errorf("Expected License: 'MIT'; got '%s'", sc.License())
	}
	authors := []string{"Mr. T", "Mr. L"}
	if !reflect.DeepEqual(sc.Authors(), authors) {
		t.Errorf("Expected Authors: '%v'; got '%v'", authors, sc.Authors())
	}
}

func Test_001_Meta_Empty(t *testing.T) {
	fname := "../../test/ekm/vdefault/001_meta_empty.ekm"
	sc, err := New(fname)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sc.Description() != defDescription {
		t.Errorf("Expected description: '%s'; got '%s'", defDescription, sc.Description())
	}
	if sc.URL() != defUrl {
		t.Errorf("Expected URL: '%s'; got '%s'", defUrl, sc.URL())
	}
	if sc.License() != defLicense {
		t.Errorf("Expected License: '%s'; got '%s'", defLicense, sc.License())
	}
	if !reflect.DeepEqual(sc.Authors(), defAuthors) {
		t.Errorf("Expected Authors: '%v'; got '%v'", defAuthors, sc.Authors())
	}
}

func Test_001_Meta_Corrupt_Values(t *testing.T) {
	fname := "../../test/ekm/vdefault/001_meta_full.ekm"
	sc, err := New(fname)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sc.Description() != "this is a test module" {
		t.Errorf("Expected description: 'this is a test module'; got '%s'", sc.Description())
	}
	sc.globals[hDescription] = starlark.MakeInt(1)
	if sc.Description() != defDescription {
		t.Errorf("Expected description: '%s'; got '%s'", defDescription, sc.Description())
	}

	authors := []string{"Mr. T", "Mr. L"}
	if !reflect.DeepEqual(sc.Authors(), authors) {
		t.Errorf("Expected Authors: '%v'; got '%v'", authors, sc.Authors())
	}
	sc.globals[hAuthors] = starlark.MakeInt(1)
	if !reflect.DeepEqual(sc.Authors(), defAuthors) {
		t.Errorf("Expected Authors: '%v'; got '%v'", defAuthors, sc.Authors())
	}
}
