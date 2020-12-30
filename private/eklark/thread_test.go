// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestThreadMustGetFilePathPanic(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Errorf("Expected error; none found")
		}
	}()

	thread := &starlark.Thread{Name: "TestThread"}
	ThreadMustGetFilePath(thread)
}
