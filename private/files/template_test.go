//Copyright (C) 2020  Germán Fuentes Capella

package files

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFromTemplate(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	err = WriteFromTemplate(tmpFile.Name(), "{{ . }}", "b")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	bdata, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	sdata := string(bdata)
	if sdata != "b" {
		t.Errorf("Expected content: '%s'; got '%s'", "b", sdata)
	}
}

func TestWriteFromTemplatePathError(t *testing.T) {
	err := WriteFromTemplate("./asdfa /adaf adfa", "{{ . }}", "b")
	if err == nil {
		t.Errorf("Expected error; found none")
	}
}

func TestWriteFromWrongTemplate(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "econkit-tmp-")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	defer os.Remove(tmpFile.Name())

	err = WriteFromTemplate(tmpFile.Name(), "{{ . ", "b")
	if err == nil {
		t.Errorf("Expected error; found none")
	}

	bdata, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}

	sdata := string(bdata)
	if sdata != "" {
		t.Errorf("Expected content: ''; got '%s'", sdata)
	}
}
