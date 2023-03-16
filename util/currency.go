package util

const (
	USD = "USD"
	EUR = "EUR"
	MDL = "MDL"
)

func IsSupporterCurrency(currency string) bool {
	switch currency {
	case USD, EUR, MDL:
		return true
	default:
		return false
	}
}
