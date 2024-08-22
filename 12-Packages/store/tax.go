package store

const DEFAULT_TAX_RATE float64 = 0.2
const MIN_THRESHOLD = 10

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = DEFAULT_TAX_RATE
	}
	if threshold < MIN_THRESHOLD {
		threshold = MIN_THRESHOLD
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) float64 {
	if product.price > taxRate.threshold {
		return product.price + (product.price * taxRate.rate)
	}
	return product.price
}