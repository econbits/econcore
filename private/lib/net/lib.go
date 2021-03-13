// Copyright (C) 2021  Germán Fuentes Capella

package net

import (
	"github.com/econbits/econkit/private/lib/net/http"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "net",
		Fns: []*slang.Fn{
			http.Fn,
		},
	}
)
