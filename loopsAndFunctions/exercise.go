package main

import (
	"fmt"
	"math"
)

// Small Diffeneces
const Small = 1.0 / (1 << 20)

// Sqrt Exported method
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		diff := (z*z - x) / (2 * z)
		if math.Abs(diff) < Small {
			break
		}
		z -= diff

		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Small)
	fmt.Println(Sqrt(2))
	fmt.Println("Actual value by math.Sqrt", math.Sqrt(2))
}
