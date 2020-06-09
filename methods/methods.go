package main

import (
	"fmt"
	"math"
)

// Vertex struct
type Vertex struct {
	X, Y float64
}

// Abs Vertex.Abs Methods
func (v Vertex) Abs() float64 {
	// (v Vertex) as receiver argument.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs Functions
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MyFloat alias of float64
type MyFloat float64

// Abs MyFloat.Abs()
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Scale *Vertext.Scale
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// Notice: Differences between Pointer and Value as receiver argument.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale Functions
func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScalePointer Functions
func ScalePointer(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
	Scale(v, 10)
	fmt.Println(Abs(v))
	ScalePointer(&v, 10)
	fmt.Println(Abs(v))

	vv := Vertex{3, 4}
	vv.Scale(2)
	ScalePointer(&vv, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScalePointer(p, 8)
	fmt.Println(vv, p)

	vvv := Vertex{3, 4}
	fmt.Println(vvv.Abs())
	fmt.Println(Abs(vvv))

	// p.Abs() is interpreted as (*p).Abs().
	pp := &Vertex{4, 3}
	fmt.Println(pp.Abs())
	fmt.Println(Abs(*pp))

	// Differences between value receiver and pointer receiver
	vvvv := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", vvvv, vvvv.Abs())
	vvvv.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", vvvv, vvvv.Abs())

}
