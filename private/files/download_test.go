//Copyright (C) 2020  Germán Fuentes Capella

package files

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	expstr := "EconKit Test"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, expstr)
	}))
	defer ts.Close()

	err = Download(ts.URL, tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	bdata, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	sdata := strings.Trim(string(bdata), "\n")
	if sdata != expstr {
		t.Errorf("Expected content: '%s'; got '%s'", expstr, sdata)
	}
}

func TestURLError(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	err = Download("", tmpFile.Name())
	if err == nil {
		t.Errorf("Expected error; none found")
	}

	bdata, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	if string(bdata) != "" {
		t.Errorf("Expected content: ''; got '%s'", string(bdata))
	}
}

func TestPathError(t *testing.T) {
	resp := "EconKit Test"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer ts.Close()

	err := Download(ts.URL, "/fasdf adf/this path does not exist")
	if err == nil {
		t.Errorf("Expected error; none found")
	}
}

func TestDownloadError(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "EconKit Error", 404)
	}))
	defer ts.Close()

	err = Download(ts.URL, tmpFile.Name())
	if err == nil {
		t.Errorf("Expected error; got none")
	}

	bdata, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	sdata := strings.Trim(string(bdata), "\n")
	if sdata != "" {
		t.Errorf("Expected content: ''; got '%s'", sdata)
	}
}
