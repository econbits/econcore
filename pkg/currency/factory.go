//Copyright (C) 2020  Germán Fuentes Capella

package currency

import (
	"fmt"
)

//go:generate go run ../../tools/isosync/main.go
var currencies = map[string]Currency{}

var emptyCurrency = Currency{}

func Get(code string) (Currency, error) {
	lcode := len(code)
	if lcode != 3 {
		return emptyCurrency, fmt.Errorf("Currency code must have 3 characters. found %d", lcode)
	}
	currency, exists := currencies[code]
	if !exists {
		return emptyCurrency, fmt.Errorf("Currency with code %s not found", code)
	}
	return currency, nil
}
