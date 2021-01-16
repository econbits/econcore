// Copyright (C) 2021  Germán Fuentes Capella

package iso

import (
	"github.com/econbits/econkit/private/lib/iso/bic"
	"github.com/econbits/econkit/private/lib/iso/country"
	"github.com/econbits/econkit/private/lib/iso/currency"
	"github.com/econbits/econkit/private/lib/iso/iban"
	"github.com/econbits/econkit/private/slang"
)

var (
	Lib = &slang.Lib{
		Name: "iso",
		Fns: []*slang.Fn{
			country.Fn,
			currency.Fn,
			bic.Fn,
			iban.Fn,
		},
	}
)
