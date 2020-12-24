// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"testing"
)

func TestInvalidErrorType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("mustTypeString did not panic on an unknown error")
		}
	}()

	unknownError.mustTypeString()
}
