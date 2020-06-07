package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	// ToBe Exported value
	ToBe bool = false
	// MaxInt  Exported value
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// default values: zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// T(v) convert data to given type
	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z1 uint = uint(f1)
	fmt.Println(x, y, z1)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	u := 3.69
	fmt.Printf("u is of type %T\n", u)

}
