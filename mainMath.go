package main

import (
	"strconv"

	//"example.com/package/npv"
	"example.com/package/tv"
	"example.com/package/irr"
)

func main() {
	println(strconv.FormatFloat(irr.TotalNpvCalc([]float64{50000, 50000}, .05), 'f', 2, 32))
	println(strconv.FormatFloat(tv.TotalTvCalc(32800000, .12,.025), 'f', 2, 64))
	println(strconv.FormatFloat(irr.IrrCalc([]float64{-10000,50000, 50000}, .05), 'f', 20, 64))
}
