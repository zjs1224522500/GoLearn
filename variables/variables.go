package main

import "fmt"

// var var_name var_type
var c, python, java bool

// Init var
var k, j int = 1, 2

// Pi const pi
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// iint default 0
	var i int
	// bool default false
	fmt.Println(i, c, python, java)

	// Can init var without type, only with given specific value
	var cpp, py, jdk = true, false, "no!"
	// short delcarations (Notice: Must be in functions)
	n := 3
	fmt.Println(k, j, n, cpp, py, jdk)

	// Constants.
	// Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Numeric constants are high-precision values.
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println(float64(1.0 / 10))
}
