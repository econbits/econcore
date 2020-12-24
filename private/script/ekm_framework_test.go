// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type errorCase struct {
	name      string
	errorType ErrorType
}

func errorTypeFromString(errstring string) ErrorType {
	for _, et := range errorTypes {
		if errstring == et.mustTypeString() {
			return et
		}
	}
	return unknownError
}

func parseErrorCase(t *testing.T, fpath string) errorCase {
	name := scriptid(fpath)
	strs := strings.SplitN(name, "_", 2)
	if len(strs) != 2 {
		t.Fatalf("filename '%s' does not follow convention: ErrorType_case_name.ekm", name)
	}
	errstring := strs[0]
	errorType := errorTypeFromString(errstring)
	if errorType == unknownError {
		t.Fatalf("Error String '%s' is not a valid error type", errstring)
	}
	return errorCase{name: name, errorType: errorType}
}

func listErrorFiles(t *testing.T, dpath string) []string {
	var files []string

	err := filepath.Walk(dpath, func(path string, info os.FileInfo, err error) error {
		if path == dpath {
			return nil
		}
		if strings.HasPrefix(scriptid(path), "OK_") {
			return nil
		}
		if !strings.HasSuffix(path, ".ekm") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	return files
}

type loadTest func(path string) error

func testErrorFile(t *testing.T, errorFile string, loadTestFn loadTest) {
	ecase := parseErrorCase(t, errorFile)
	err := loadTestFn(errorFile)
	if err == nil {
		t.Fatalf("[%s] Expected %v; none found", errorFile, ecase.errorType.mustTypeString())
	}
	scriptErr, ok := err.(ScriptError)
	if !ok {
		t.Fatalf("[%s] Expected ScriptError; found: %T with value '%v'", errorFile, err, err)
	}
	if scriptErr.errorType != ecase.errorType {
		t.Fatalf(
			"[%s] Expected Error Type '%s'; found '%s' (Error msg: '%v'",
			errorFile,
			ecase.errorType.mustTypeString(),
			scriptErr.errorType.mustTypeString(),
			scriptErr,
		)
	}
	if len(err.Error()) == 0 {
		t.Fatalf(
			"[%s] Empty Error Message",
			errorFile,
		)
	}
}

func testErrorFiles(t *testing.T, rootPath string, loadTestFn loadTest) {
	errorFiles := listErrorFiles(t, rootPath)
	for _, errorFile := range errorFiles {
		t.Run(errorFile, func(t *testing.T) {
			testErrorFile(t, errorFile, loadTestFn)
		})
	}
}
