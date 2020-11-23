//Copyright (C) 2020  Germán Fuentes Capella

package money

// Compares two Money structs, by amount and currency
func (m1 Money) IsEqual(m2 Money) bool {
	if m1.amount != m2.amount {
		return false
	}
	return m1.currency.IsEqual(m2.currency)
}
