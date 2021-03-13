// Copyright (C) 2021  Germán Fuentes Capella

package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestAssertHTTPResponse(t *testing.T) {
	intv := starlark.MakeInt(1)
	_, err := AssertHTTPResponse(intv)
	if err == nil {
		t.Error("1 is not an HTTPResponse")
	}

	eUrl, _ := url.Parse("https://econbits.org")
	hresp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header:     http.Header{},
		Request: &http.Request{
			Method: "GET",
			URL:    eUrl,
		},
		Body: ioutil.NopCloser(strings.NewReader("OK")),
	}
	resp, err := newHTTPResponse(hresp)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	newresp, err := AssertHTTPResponse(resp)
	if err != nil {
		t.Error("an HTTPResponse is not identified as such")
	}
	if newresp != resp {
		t.Fatalf("expected %v; got %v", resp, newresp)
	}
}

func TestAssertHeaderDict(t *testing.T) {
	intv := starlark.MakeInt(1)
	_, err := AssertHeaderDict(intv)
	if err == nil {
		t.Error("1 is not a dict")
	}

	dict := starlark.NewDict(10)
	_, err = AssertHeaderDict(dict)
	if err != nil {
		t.Error("an empty dict is a valid dict")
	}

	err = dict.SetKey(intv, intv)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	_, err = AssertHeaderDict(dict)
	if err == nil {
		t.Error("error: the header dict contains a numeric key")
	}
	err = dict.Clear()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	str := starlark.String("abc")
	err = dict.SetKey(str, str)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	_, err = AssertHeaderDict(dict)
	if err == nil {
		t.Error("error: the header dict contains a string value, instead of string list value")
	}
	err = dict.Clear()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	list := starlark.NewList([]starlark.Value{str})
	err = dict.SetKey(str, list)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	_, err = AssertHeaderDict(dict)
	if err != nil {
		t.Error("it is a valid dict: map[string][]string")
	}

	list = starlark.NewList([]starlark.Value{intv})
	err = dict.SetKey(str, list)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	_, err = AssertHeaderDict(dict)
	if err == nil {
		t.Error("error: the header dict contains a int list value, instead of string list value")
	}
}
