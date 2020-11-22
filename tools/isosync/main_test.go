//Copyright (C) 2020  Germán Fuentes Capella
package main

// TODO these tests are very rudimentary. Rework them to increase test coverage
// and testability of the code

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	err := os.Remove(dataPath)
	if err != nil {
		t.Errorf("Unexpected error deleting the file: %v", err)
	}
	defer func() {
		if e := recover(); e != nil {
			t.Errorf("Unexpected error downloading the file: %v", e)
		}
	}()
	downloadISO4217XML()
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		t.Errorf("Unexpected error reading the file content: %v", err)
	}
	strdata := string(data)
	if !strings.Contains(strdata, "<ISO_4217") {
		t.Errorf("Data file does not contain start targ: <ISO_4217>")
	}
	if !strings.Contains(strdata, "</ISO_4217") {
		t.Errorf("Data file does not contain end targ: </ISO_4217>")
	}
}

func TestLoadCurrencyWithoutUnits(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Errorf("Unexpected error loading currencies: %v", e)
		}
	}()
	currencies := loadCurrencies()
	na_currency, exists := currencies["XBA"]
	if !exists {
		t.Errorf("XBA not found in ISO4217 currencies")
	}
	if na_currency.Units != 0 {
		t.Errorf("Expected XBA units: 0; got: %d", na_currency.Units)
	}
}

func TestFullFlow(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Errorf("Unexpected error creating Go file: %v", e)
		}
	}()
	err := os.Remove(goPath)
	if err != nil {
		t.Errorf("Unexpected error deleting the Go file: %v", err)
	}

	main()

	_, err = os.Stat(goPath)
	if err != nil {
		t.Errorf("Unexpected error checking Go file: %v", err)
	}
}
