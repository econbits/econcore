// Copyright (C) 2021  Germán Fuentes Capella

package ekm

import (
	"fmt"

	"github.com/econbits/econkit/private/lib/account"
	"github.com/econbits/econkit/private/lib/auth"
	"github.com/econbits/econkit/private/lib/datetime"
	"github.com/econbits/econkit/private/lib/fin"
	"github.com/econbits/econkit/private/lib/iso"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

func load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	libs := []*slang.Lib{
		account.Lib,
		auth.Lib,
		datetime.Lib,
		fin.Lib,
		iso.Lib,
	}
	for _, lib := range libs {
		if module == lib.Name {
			return lib.Load(), nil
		}
	}
	return nil, fmt.Errorf("unknown module: %s", module)
}
