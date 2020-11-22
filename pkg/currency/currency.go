//Copyright (C) 2020  Germán Fuentes Capella

package currency

// Currency as defined in ISO 4217.
// Note: the list of countries in which the currency is used is not parsed.
type Currency struct {
	name  string
	code  string
	id    uint32
	units uint8
}

func (c Currency) Name() string {
	return c.name
}

func (c Currency) Code() string {
	return c.code
}

func (c Currency) Id() uint32 {
	return c.id
}

func (c Currency) Units() uint8 {
	return c.units
}
