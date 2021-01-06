// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestPreProcessIbanKindError(t *testing.T) {
	_, err := preprocessIbanKind(starlark.MakeInt(1))
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}

func TestPreProcessWalletKindError(t *testing.T) {
	_, err := preprocessWalletKind(starlark.MakeInt(1))
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}
