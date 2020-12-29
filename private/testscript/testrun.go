// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/econbits/econkit/private/eklark"
)

func isScript(dirPath, filePath string) bool {
	if filePath == dirPath {
		return false
	}
	return strings.HasSuffix(filePath, ".ekm")
}

func testRunner(filePath string, testFn TestFn) *eklark.EKError {
	testCase := ParseTestCase(filePath)
	if testCase.GotError != nil {
		return testCase.GotError
	}
	Run(testCase, testFn)
	return testCase.GotError
}

type FailFn func(t *testing.T, err error)

func Fail(t *testing.T, err error) {
	t.Fatal(err.Error())
}

func TestRun(t *testing.T, dirPath string, testFn TestFn, failFn func(t *testing.T, err error)) {
	filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if isScript(dirPath, filePath) {
		}
		t.Run(filePath, func(t *testing.T) {
			err := testRunner(filePath, testFn)
			if err != nil {
				failFn(t, err)
			}
		})
		return nil
	})
}
