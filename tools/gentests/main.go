// Copyright (C) 2021  Germán Fuentes Capella

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/ekm"
	"github.com/econbits/econkit/private/lib/universe"
	"go.starlark.net/starlark"
)

const (
	testPath = "test/ekm/vdefault"
	outPath  = "test/outfiles"
)

func content(path string, in []byte, err error) []byte {
	buf := bytes.NewBufferString("FILE: " + path + "\n\nSCRIPT:\n\n")
	buf.Write(in)
	if err != nil {

		var ekerr *ekerrors.EKError
		if errors.As(err, &ekerr) {
			buf.WriteString("\n\nBACKTRACE:\n\n" + ekerr.Backtrace())
		}

		buf.WriteString("\n\nERRORS:\n\n")
		i := 0
		for {

			buf.WriteString(
				strings.Repeat("\t", i) + fmt.Sprintf(" [%T] ", err) + err.Error() + "\n",
			)

			err = errors.Unwrap(err)
			if err == nil {
				break
			}

			i++
		}
	} else {
		buf.WriteString("ERRORS:\n\nNone\n")
	}
	return buf.Bytes()
}

func write(path string, scripterr error) error {
	inscript, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	bytes := content(path, inscript, scripterr)

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
	fmt.Println(opath)
	odir := filepath.Dir(opath)
	err = os.MkdirAll(odir, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(opath, bytes, 0644)
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
