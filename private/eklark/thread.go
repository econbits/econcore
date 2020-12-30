// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"go.starlark.net/starlark"
)

const (
	localFilePath = "FilePath"
)

func NewThread(fpath string) *starlark.Thread {
	name := ScriptId(fpath)
	t := &starlark.Thread{Name: name}
	t.SetLocal(localFilePath, fpath)
	return t
}

func ThreadMustGetFilePath(t *starlark.Thread) string {
	ifce := t.Local(localFilePath)
	str, ok := ifce.(string)
	if !ok {
		panic("Thread was not initialized with eklark.NewThread()")
	}
	return str
}
