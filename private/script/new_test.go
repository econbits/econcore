//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"
)

func TestNewNoFile(t *testing.T) {
	fpath := "this_file_does_not_exist.star"
	_, err := New(fpath)
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}
