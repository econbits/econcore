// Copyright (C) 2021  Germán Fuentes Capella

package account

import (
	"testing"

	"go.starlark.net/starlark"
)

func TestMatchKindError(t *testing.T) {
	_, err := matchAccountKind(starlark.MakeInt(1))
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}

func TestMatchKindSuccess(t *testing.T) {
	_, err := matchAccountKind(starlark.String("checking"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMatchMissingKindError(t *testing.T) {
	_, err := matchAccountKind(starlark.String("not-checking"))
	if err == nil {
		t.Fatalf("expected error; none found")
	}
}
