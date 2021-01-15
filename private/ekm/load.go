// Copyright (C) 2021  Germán Fuentes Capella

package ekm

import (
	"fmt"

	"github.com/econbits/econkit/private/lib/datetime"
	"github.com/econbits/econkit/private/lib/iso"
	"go.starlark.net/starlark"
)

func load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	if module == datetime.Lib.Name {
		return datetime.Lib.Load(), nil
	}
	if module == iso.Lib.Name {
		return iso.Lib.Load(), nil
	}
	return nil, fmt.Errorf("unknown module: %s", module)
}
