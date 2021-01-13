// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"reflect"
	"testing"

	"go.starlark.net/starlark"
)

func Test2IntRange(t *testing.T) {
	want := []int{1, 2}
	slist := IRtoVR(want)

	got, err := VRtoIR(slist)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Error: want %v != got %v", want, got)
	}
}

func TestVRtoIRError(t *testing.T) {
	s := starlark.String("hello")
	_, err := VRtoIR([]starlark.Value{s})
	if err == nil {
		t.Errorf("Expected error, none found")
	}
}
