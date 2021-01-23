// Copyright (C) 2021  Germán Fuentes Capella

package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/econbits/econkit/private/ekm"
	"github.com/econbits/econkit/private/lib/universe"
	"go.starlark.net/starlark"
)

const (
	testPath = "test/ekm/vdefault"
	outPath  = "test/outfiles"
)

func write(path string, scripterr error) error {
	inscript, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	buf := bytes.NewBufferString("FILE: " + path + "\n\nSCRIPT:\n\n")
	buf.Write(inscript)
	buf.WriteString("\n\nERRORS:\n\n")
	if scripterr != nil {
		err = scripterr
		i := 0
		for {
			err = errors.Unwrap(err)
			if err != nil {
				buf.WriteString(
					strings.Repeat(".", i) + err.Error() + "\n",
				)
				i++
			} else {
				break
			}
		}
	} else {
		buf.WriteString("None\n")
	}

	opath, err := filepath.Abs(
		strings.Replace(
			strings.Replace(path, ".ekm", ".txt", 1),
			testPath,
			outPath,
			1,
		),
	)
	if err != nil {
		return err
	}
	odir := filepath.Dir(opath)
	err = os.MkdirAll(odir, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(opath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := filepath.Walk(testPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(path, ".ekm") {
			return nil
		}
		if strings.Contains(path, "/000_smalltests/") {
			thread := &starlark.Thread{Name: path, Load: ekm.Load}
			_, err = starlark.ExecFile(thread, path, nil, universe.Lib.Load())
			return write(path, err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
