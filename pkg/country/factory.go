//Copyright (C) 2020  Germán Fuentes Capella

package country

import (
	"fmt"
)

//go:generate go run ../../tools/synciso3166/main.go
var countries = map[string]Country{}

func Get(alpha2 string) (Country, error) {
	l2 := len(alpha2)
	if l2 != 2 {
		return Country{}, fmt.Errorf("Country alpha2 must have 2 characters. found %d", l2)
	}
	country, exists := countries[alpha2]
	if !exists {
		return Country{}, fmt.Errorf("Country with alpha2 code %s not found", alpha2)
	}
	return country, nil
}

func MustGet(alpha2 string) Country {
	c, err := Get(alpha2)
	if err != nil {
		panic(err.Error())
	}
	return c
}
