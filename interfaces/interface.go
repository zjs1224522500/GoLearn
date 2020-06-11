package main

import (
	"fmt"
	"math"
)

// Abser Interface
type Abser interface {
	Abs() float64
}

// MyFloat type
type MyFloat float64

// Abs Implement method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex type
type Vertex struct {
	X, Y float64
}

// Abs Implement method
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// I interface
type I interface {
	M()
}

// T struct
type T struct {
	S string
}

// M Implement I
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// TT struct
type TT struct {
	S string
}

// M Implement
func (tt *TT) M() {
	// nil value receiver
	// handler nil value
	if tt == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(tt.S)
}

// F type
type F float64

// M implement method
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func testEmptyInterfaces() {
	var i interface{}
	describeEmptyInterface(i)

	i = 42
	describeEmptyInterface(i)

	i = "hello"
	describeEmptyInterface(i)
}

func interfaceValueTest() {
	var i I
	// runtime exception, because there is no type inside the interface tuple to indicate which concrete method to call.
	// describe(i)
	// i.M()

	var t *TT
	i = t
	describe(i)
	i.M()

	i = T{"world"}
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

func implicitInterfaceTest() {
	var i I = T{"hello"}
	i.M()
}

func basicInterfaceTest() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

func main() {
	basicInterfaceTest()
	implicitInterfaceTest()
	interfaceValueTest()
	testEmptyInterfaces()
}
