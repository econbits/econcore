// Copyright (C) 2021  Germán Fuentes Capella

package encoding

import (
	"io/ioutil"
	"testing"
)

func TestLib(t *testing.T) {
	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	dirs := 0
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			dirs += 1
		}
	}
	dirs = 2 * dirs // *_encode & *_decode
	if len(Lib.Fns) != dirs {
		t.Fatalf("expected %d Fns; got %v", dirs, Lib.Fns)
	}
}
