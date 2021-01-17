// Copyright (C) 2021  Germán Fuentes Capella

package ibanaccount

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestMatchIbanKindError(t *testing.T) {
	_, err := matchIbanKind(starlark.MakeInt(1))
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}
