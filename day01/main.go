package main

import (
	"fmt"
	"math"
)

var inputs = [...]float64{
	50669,
	83199,
	108957,
	102490,
	121216,
	57848,
	76120,
	121042,
	143774,
	106490,
	76671,
	119551,
	109598,
	124949,
	148026,
	119862,
	112854,
	96289,
	73573,
	142714,
	109875,
	126588,
	69748,
	62766,
	104446,
	61766,
	133199,
	118306,
	141410,
	64420,
	129370,
	69444,
	104178,
	109696,
	54654,
	126517,
	138265,
	87100,
	130934,
	138946,
	118078,
	135109,
	137330,
	130913,
	50681,
	53071,
	63070,
	144616,
	64122,
	122603,
	86612,
	144666,
	62572,
	72247,
	79005,
	102223,
	82585,
	54975,
	68539,
	107964,
	148333,
	100269,
	134101,
	115960,
	75866,
	99371,
	56685,
	142199,
	102107,
	84075,
	112733,
	92180,
	131056,
	142803,
	145495,
	70900,
	73129,
	60764,
	77576,
	99160,
	61897,
	78675,
	100890,
	74787,
	131911,
	93106,
	91267,
	142663,
	130649,
	103857,
	81178,
	78896,
	73745,
	84463,
	92926,
	121715,
	74959,
	71911,
	89102,
	53428,
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take
// its mass, divide by three, round down, and subtract 2.
//
// f = floor(m / 3) - 2
func calculateFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func main() {
	var runningSum float64 = 0
	for _, num := range inputs {
		runningSum += calculateFuel(num)
	}
	fmt.Printf("%f\n", runningSum)
}