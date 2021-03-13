// Copyright (C) 2021  Germán Fuentes Capella

package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

type HTTPResponse struct {
	slang.EKValue
}

const (
	FnName      = "http"
	typeName    = "HTTPResponse"
	fStatusCode = "status_code"
	fHeaders    = "headers"
	fUrl        = "url"
	fStatus     = "status"
	fText       = "text"
	// TODO cookies
)

var (
	Fn = &slang.Fn{
		Name:     FnName,
		Callback: httpFn,
	}
)

func validate(method, url starlark.String) error {
	methodstr := string(method)
	if methodstr != "GET" && methodstr != "POST" {
		return errors.New(fmt.Sprintf("HTTP methods supported are GET and POST; found: %v", method))
	}
	if !strings.HasPrefix(string(url), "https://") {
		return errors.New(fmt.Sprintf("Expected HTTPS; found: %s", url))
	}
	return nil
}

func newHTTPResponse(hresp *http.Response) (*HTTPResponse, error) {
	headerD := starlark.NewDict(len(hresp.Header))
	for headerk, headervl := range hresp.Header {
		headerlist := starlark.NewList([]starlark.Value{})
		for _, headerv := range headervl {
			headerlist.Append(starlark.String(headerv))
		}
		err := headerD.SetKey(starlark.String(headerk), headerlist)
		if err != nil {
			return nil, err
		}
	}
	bytes, err := io.ReadAll(hresp.Body)
	if err != nil {
		return nil, err
	}
	text := starlark.String(bytes)
	return &HTTPResponse{
		slang.NewEKValue(
			typeName,
			[]string{
				fStatus,
				fStatusCode,
				fHeaders,
				fText,
				fUrl,
			},
			map[string]starlark.Value{
				fStatus:     starlark.String(hresp.Status),
				fStatusCode: starlark.MakeInt(hresp.StatusCode),
				fHeaders:    headerD,
				fText:       text,
				fUrl:        starlark.String(hresp.Request.URL.String()),
			},
			map[string]slang.PreProcessFn{
				fStatus:     slang.AssertString,
				fStatusCode: slang.AssertInt,
				fHeaders:    AssertHeaderDict,
				fText:       slang.AssertString,
				fUrl:        slang.AssertString,
			},
			slang.NoMaskFn,
		),
	}, nil
}

func httpFn(
	thread *starlark.Thread,
	builtin *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {
	var method, url starlark.String
	err := starlark.UnpackArgs(
		builtin.Name(), args, kwargs,
		"method", &method,
		"url", &url,
	)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	err = validate(method, url)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	ctx := slang.LocalContext(thread)
	req, err := http.NewRequestWithContext(ctx, string(method), string(url), nil)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, ekerrors.Wrap(
			errorClass,
			err,
			[]ekerrors.Format{},
		)
	}
	defer resp.Body.Close()
	return newHTTPResponse(resp)
}
