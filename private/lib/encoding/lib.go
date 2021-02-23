// Copyright (C) 2021  Germán Fuentes Capella

package encoding

import (
	"github.com/econbits/econkit/private/lib/encoding/json"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "encoding",
		Fns: []*slang.Fn{
			json.DecFn,
		},
	}
)
