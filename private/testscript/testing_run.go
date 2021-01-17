// Copyright (C) 2020  Germán Fuentes Capella

package testscript

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func isScript(dirPath, filePath string) bool {
	if filePath == dirPath {
		return false
	}
	return strings.HasSuffix(filePath, ".ekm")
}

func testRunner(filePath string, epilogue starlark.StringDict, load LoadFn, testFn TestFn) error {
	testCase := ParseTestCase(filePath)
	if testCase.AbortError != nil {
		return testCase.AbortError
	}
	RunTestCase(testCase, epilogue, load, testFn)
	return testCase.AbortError
}

type FailFn func(t *testing.T, err error)

func Fail(t *testing.T, err error) {
	t.Fatal(err.Error())
}

func resolvePath(dirPath string) (string, error) {
	_, err := os.Stat(dirPath)
	if err != nil {
		dirPath = filepath.Join("test", "ekm", "vdefault", dirPath)
		// in this case, we try to find the full path
		for i := 0; i < 5; i++ {
			dirPath = filepath.Join("..", dirPath)
			_, err := os.Stat(dirPath)
			if err == nil {
				return dirPath, nil
			}
		}
		return "", err
	}
	return dirPath, nil
}

func TestingRun(
	t *testing.T,
	dirPath string,
	epilogue starlark.StringDict,
	load LoadFn,
	testFn TestFn,
	failFn func(t *testing.T, err error),
) {
	dirPath, err := resolvePath(dirPath)
	if err != nil {
		failFn(t, err)
	}
	filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if isScript(dirPath, filePath) {
			t.Run(filePath, func(t *testing.T) {
				err := testRunner(filePath, epilogue, load, testFn)
				if err != nil {
					failFn(t, err)
				}
			})
		}
		return nil
	})
}
