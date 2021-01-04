// Copyright (C) 2020  Germán Fuentes Capella

package datetime

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertDateTime(t *testing.T) {
	intv := starlark.MakeInt(1)
	_, err := AssertDateTime(intv)
	if err == nil {
		t.Error("1 is not a date")
	}

	expectstr := "2020-01-30"
	dt, err := New("2006-01-02", expectstr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	newdt, err := AssertDateTime(dt)
	if err != nil {
		t.Error("a date is not identified a date")
	}
	if newdt != dt {
		t.Fatalf("expected %v; got %v", dt, newdt)
	}

	gotstr := dt.String()
	if gotstr != expectstr {
		t.Fatalf("expected %v; got %v", expectstr, gotstr)
	}
}
