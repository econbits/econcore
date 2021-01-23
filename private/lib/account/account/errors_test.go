// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"
)

func TestFormatError(t *testing.T) {
	msg := "test"
	got := FormatError(msg)
	if msg != got {
		t.Fatalf("msg(%s) != got(%s)", msg, got)
	}
}
