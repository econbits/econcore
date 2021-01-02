// Copyright (C) 2020  Germán Fuentes Capella

package datetime

import (
	"testing"
	"time"

	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func TestScripts(t *testing.T) {
	dpath := "../../../test/ekm/vdefault/000_smalltests/ekres/datetime/"
	fn := DateTimeFn
	epilogue := starlark.StringDict{fn.Name: fn.Builtin()}
	testscript.TestingRun(t, dpath, epilogue, testscript.ExecScriptFn, testscript.Fail)
}

func TestAssertDateTime(t *testing.T) {
	intv := starlark.MakeInt(1)
	err := AssertDateTime(intv)
	if err == nil {
		t.Error("1 is not a date")
	}

	expectstr := "2020-01-30"
	dt, err := New("2006-01-02", expectstr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	err = AssertDateTime(dt)
	if err != nil {
		t.Error("a date is not identified a date")
	}

	gotstr := dt.String()
	if gotstr != expectstr {
		t.Fatalf("expected %v; got %v", expectstr, gotstr)
	}
}

func TestNewFromTime(t *testing.T) {
	expectdate, err := time.Parse("2006-01-02", "2020-01-30")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	dt := NewFromTime(expectdate)
	gotdate := dt.Time()
	if !expectdate.Equal(gotdate) {
		t.Fatalf("expected %v; got %v", expectdate, gotdate)
	}
}
