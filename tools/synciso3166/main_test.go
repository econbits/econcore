//Copyright (C) 2020  Germán Fuentes Capella

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLoadCountries(t *testing.T) {
	countries, err := load(dataPath)
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	ar_country, exists := countries["AR"]
	if !exists {
		t.Errorf("Argentina not found in ISO3166 countries")
	}
	if ar_country.Name != "ARGENTINA" {
		t.Errorf("Expected AR name: ARGENTINA; got: %s", ar_country.Name)
	}
}

func TestLoadCountryWithWrongPath(t *testing.T) {
	_, err := load("./adfa/asdfadf adfad")
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}

func TestLoadCountryWithWrongContent(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, strings.NewReader(`Abc,Def
x,y,z,0,9`))
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	_, err = load(tmpFile.Name())
	if err == nil {
		t.Errorf("Expected error; got none")
	}

	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	_, err = io.Copy(tmpFile, strings.NewReader(`Abc,Def,Ghi
jkl,mno,pqr
`))
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	_, err = load(tmpFile.Name())
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}

func TestFullCountryGen(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Errorf("Unexpected error creating Go file: %v", e)
		}
	}()

	bdata, err := ioutil.ReadFile(dataPath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(bdata))
	}))
	defer ts.Close()

	initialUrl := iso3166Url
	iso3166Url = ts.URL
	defer func() { iso3166Url = initialUrl }()

	err = os.Remove(goPath)
	if err != nil {
		t.Errorf("Unexpected error deleting the Go file: %v", err)
	}

	main()

	_, err = os.Stat(goPath)
	if err != nil {
		t.Errorf("Unexpected error checking Go file: %v", err)
	}
}
