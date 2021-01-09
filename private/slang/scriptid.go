// Copyright (C) 2020  Germán Fuentes Capella

package slang

import (
	"path/filepath"
)

func ScriptId(fpath string) string {
	name := filepath.Base(fpath)
	ext := filepath.Ext(name)
	return name[0 : len(name)-len(ext)]
}
