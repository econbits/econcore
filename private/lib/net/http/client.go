// Copyright (C) 2021  Germán Fuentes Capella

package http

import (
	"net/http"
)

var (
	httpClient = &http.Client{}
)

type RoundTripFn func(req *http.Request) (*http.Response, error)

type fTransport struct {
	roundTripFn RoundTripFn
}

func (tr *fTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return tr.roundTripFn(req)
}

func InjectInClient(roundTripFn RoundTripFn) {
	tr := &fTransport{
		roundTripFn: roundTripFn,
	}
	httpClient = &http.Client{Transport: tr}
}

func ResetClient() {
	httpClient = &http.Client{}
}
