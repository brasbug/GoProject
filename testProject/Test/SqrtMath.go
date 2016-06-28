package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for {
		y := z - (z*z - x) / (2 * z)
		if math.Abs(y - z) < 1e-14 {
			return y
		}
		z = y;
	}
	return z
}



func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}