// Copyright (C) 2021  Germán Fuentes Capella

package universe

import (
	"github.com/econbits/econkit/private/lib/universe/assert"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "universe",
		Fns: []*slang.Fn{
			assert.AssertFn,
		},
	}
)
