//Copyright (C) 2020  Germán Fuentes Capella

package country

// Country per ISO 3166
type Country struct {
	name   string
	alpha2 string
}

func (c Country) Name() string {
	return c.name
}

func (c Country) Alpha2() string {
	return c.alpha2
}
