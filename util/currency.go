package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	THB = "THB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, THB:
		return true
	}
	return false
}
