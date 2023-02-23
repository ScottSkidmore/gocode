package irr

import (
	"math"
)

func IrrCalc(cashFlow []float64, discountRate float64) float64 {
		// Set initial guess
		guess := discountRate+5
		
		for i := 0; i < 100; i++ {
			npv := 0.0
			dnpv := 0.0
			
			for f, cash := range cashFlow {
				npv += cash/ math.Pow(1+guess, float64(f))
				dnpv -= float64(f) * cash / math.Pow(1+guess, float64(f+1))
			}
			
			guess -= npv / dnpv
			
			if math.Abs(npv) < 1e-8 {
				return guess
			}
		}
		
		return math.NaN()
	}

func TotalNpvCalc(cashFlow []float64, discountRate float64) float64 {
	total := float64(0)
	for i, v := range cashFlow {
		total = total + (v / float64(math.Pow(float64(1+discountRate), float64(i+1))))
		total = math.Round(total*100) / 100
	}
	return total
}