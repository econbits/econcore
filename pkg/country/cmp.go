//Copyright (C) 2020  Germán Fuentes Capella

package country

// Compares 2 Countries by Alpha-2 code
func (c1 Country) IsEqual(c2 Country) bool {
	return c1.alpha2 == c2.alpha2
}
