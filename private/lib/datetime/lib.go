// Copyright (C) 2021  Germán Fuentes Capella

package datetime

import (
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "datetime",
		Fns: []*slang.Fn{
			datetime.Fn,
		},
	}
)
