// Copyright (C) 2021  Germán Fuentes Capella

package slang

import (
	"context"

	"go.starlark.net/starlark"
)

const (
	localCtx = "ctx"
)

func SetLocalContext(thread *starlark.Thread, ctx context.Context) {
	thread.SetLocal(localCtx, ctx)
}

func LocalContext(thread *starlark.Thread) context.Context {
	ictx := thread.Local(localCtx)
	ctx, _ := ictx.(context.Context)
	return ctx
}
