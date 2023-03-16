package store

const defaultTaxRate float64 = 0.2
const minTreshold = 10

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minTreshold {
		threshold = minTreshold
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(p *Product) float64 {
	if p.price > taxRate.threshold {
		return p.price + p.price*taxRate.rate
	}
	return p.price
}
