// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"testing"

	"github.com/econbits/econkit/private/eklark"
)

func TestIsScript(t *testing.T) {
	if isScript("/tmp", "/tmp") {
		t.Fatal("/tmp is not a script")
	}
	if isScript("/tmp", "filename.txt") {
		t.Fatal("filename.txt is not a script")
	}
	if !isScript("/tmp", "script.ekm") {
		t.Fatal("script.ekm is a script")
	}
}

func TestTestRunner(t *testing.T) {
	err := testRunner("OK_script.ekm", func(path string) *eklark.EKError {
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestTestRunnerOnErrorFile(t *testing.T) {
	err := testRunner("script.ekm", func(path string) *eklark.EKError {
		return nil
	})
	if err == nil {
		t.Fatal("Expected error; none found")
	}
}
