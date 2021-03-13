// Copyright (C) 2021  Germán Fuentes Capella

package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/testscript"
	"go.starlark.net/starlark"
)

func roundTripFn(req *http.Request) (*http.Response, error) {
	if req.Method != "GET" && req.Method != "POST" {
		panic(fmt.Sprintf("HTTP Request Method should have been validated; found %s", req.Method))
	}
	if req.URL.Host == "econkit.org" {
		if req.URL.Path == "/ok" {
			header := http.Header{}
			header.Add("key", "value")
			resp := &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Header:     header,
				Request:    req,
				Body:       ioutil.NopCloser(strings.NewReader("OK")),
			}
			return resp, nil
		}
	}
	if req.URL.Host == "this_domain_does_not_exist" {
		return nil, fmt.Errorf("invalid domain 'this_domain_does_not_exist'")
	}
	panic(fmt.Sprintf("Test case not covered for %v (host: %s; path: %s)", req, req.URL.Host, req.URL.Path))
}

func TestScripts(t *testing.T) {
	InjectInClient(roundTripFn)
	defer ResetClient()

	dpath := "000_smalltests/lib/net/http/"
	testscript.TestingRun(
		t,
		dpath,
		universe.Lib.Load(),
		func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			if module == "net" {
				sd := starlark.StringDict{
					Fn.Name: Fn.Builtin(),
				}
				return sd, nil
			}
			return nil, fmt.Errorf("unknown module: %s", module)
		},
		testscript.ExecScriptFn,
		testscript.Fail,
	)
}
