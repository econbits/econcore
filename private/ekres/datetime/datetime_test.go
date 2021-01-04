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
