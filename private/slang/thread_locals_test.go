// Copyright (C) 2021  Germán Fuentes Capella

package slang

import (
	"context"
	"testing"

	"go.starlark.net/starlark"
)

func TestLocalContext(t *testing.T) {
	thread := &starlark.Thread{}
	ctx := context.Background()
	SetLocalContext(thread, ctx)

	octx := LocalContext(thread)
	if ctx != octx {
		t.Errorf("%v != %v", ctx, octx)
	}
}
