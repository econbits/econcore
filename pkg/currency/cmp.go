//Copyright (C) 2020  Germán Fuentes Capella

package currency

// Compares 2 currencies, by code
func (c1 Currency) IsEqual(c2 Currency) bool {
	return c1.code == c2.code
}
