// Copyright (C) 2020  Germán Fuentes Capella

package date

import (
	"testing"
	"time"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/date/"
	epilogue := starlark.StringDict{"date": DateFn.Builtin()}
	testscript.TestingRun(t, dpath, epilogue, testscript.ExecScriptFn, testscript.Fail)
}

func TestAssertDate(t *testing.T) {
	intv := starlark.MakeInt(1)
	err := AssertDate(intv)
	if err == nil {
		t.Error("1 is not a date")
	}
	datev := New("2006-01-02", "2020-01-30")
	err = AssertDate(datev)
	if err != nil {
		t.Error("a date is not identified a date")
	}
}

func TestNewFromTime(t *testing.T) {
	expectdate, err := time.Parse("2006-01-02", "2020-01-30")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	datev := NewFromTime(expectdate)
	gotdate, err := datev.Time()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if !expectdate.Equal(gotdate) {
		t.Fatalf("expected %v; got %v", expectdate, gotdate)
	}
}
