package tv

import (
	"math"
)


func TotalTvCalc(freeCashFlow float64, discountRate float64, growthRate float64) float64 {
	total:=(freeCashFlow*(1+growthRate))/(discountRate-growthRate)
	total = math.Round(total*100) / 100
	return total
}