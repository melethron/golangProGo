package main

type Service struct {
	description    string
	durationMonths int8
	monthlyFee     float64
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return float64(s.durationMonths) * s.monthlyFee
	}
	return s.monthlyFee
}
