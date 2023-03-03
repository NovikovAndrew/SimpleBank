package util

// const for all suported currency
const (
	USD = "USD"
	EUR = "EUR"
	KZT = "KZT"
	RUB = "RUB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, KZT, RUB:
		return true
	}
	return false
}
