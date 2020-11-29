//Copyright (C) 2020  Germán Fuentes Capella
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLoadCurrencyWithoutUnits(t *testing.T) {
	currencies, err := load(dataPath)
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	na_currency, exists := currencies["XBA"]
	if !exists {
		t.Errorf("XBA not found in ISO4217 currencies")
	}
	if na_currency.Units != 0 {
		t.Errorf("Expected XBA units: 0; got: %d", na_currency.Units)
	}
}

func TestLoadCurrencyWithWrongPath(t *testing.T) {
	_, err := load("./adfa/asdfadf adfad")
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}

func TestLoadCurrencyWithWrongContent(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = load(tmpFile.Name())
	if err == nil {
		t.Errorf("Expected error; got none")
	}
}

func TestFullFlow(t *testing.T) {
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

	initialUrl := iso4217Url
	iso4217Url = ts.URL
	defer func() { iso4217Url = initialUrl }()

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
