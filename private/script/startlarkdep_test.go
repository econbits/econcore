// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

type deferFunc func()

func script2file(script string) (*os.File, deferFunc, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(tmpFile, strings.NewReader(script))
	if err != nil {
		return nil, nil, err
	}

	return tmpFile, func() { os.Remove(tmpFile.Name()) }, nil
}

var (
	scriptFibonacci = `
def fibonacci(n):
	res = list(range(n))
	for i in res[2:]:
		res[i] = res[i-2] + res[i-1]
	return res
`
)

func TestStarlark(t *testing.T) {
	scriptFile, deferfunc, err := script2file(scriptFibonacci)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}
	defer deferfunc()

	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "Test Startlark Dependency"}
	globals, err := starlark.ExecFile(thread, scriptFile.Name(), nil, nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}

	// Retrieve a module global.
	fibonacci := globals["fibonacci"]
	if fibonacci == nil {
		t.Errorf("Function expected, nil found")
		return
	}

	// Call Starlark function from Go.
	v, err := starlark.Call(thread, fibonacci, starlark.Tuple{starlark.MakeInt(10)}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}

	got, err := VtoIR(v)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}

	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v != Got %v", expected, got)
	}
}
