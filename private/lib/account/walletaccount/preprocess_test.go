// Copyright (C) 2021  Germán Fuentes Capella

package walletaccount

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestMatchWalletKindError(t *testing.T) {
	_, err := matchWalletKind(starlark.MakeInt(1))
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}
