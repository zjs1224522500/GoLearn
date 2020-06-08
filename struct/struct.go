package main

import (
	"fmt"
	"strings"
)

func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(&p) // read the address of pointer p
	fmt.Println(p)  // read the addrees of i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// Vertex Exported Struct
type Vertex struct {
	X int
	Y int
}

func structTest() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)
	var (
		v1      = Vertex{1, 2}  // has type Vertex
		v2      = Vertex{X: 1}  // Y:0 is implicit
		v3      = Vertex{}      // X:0 and Y:0
		pointer = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, pointer, v2, v3)

}

// Arrays: Fixed size
func arrayTest() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slice: Dynamic size
// Not store the data, like the references of array.
func sliceTest() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// no.1 - no.3
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

}

func sliceDefaultIndexTest() {

	// slice default inde value
	sss := []int{2, 3, 5, 7, 11, 13}

	sss = sss[1:4]
	fmt.Println(sss)

	sss = sss[:2]
	fmt.Println(sss)

	sss = sss[1:]
	fmt.Println(sss)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSliceTest() {
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5) // len(a)=5
	printSliceWithName("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSliceWithName("b", b)

	c := b[:2]
	printSliceWithName("c", c)

	d := c[2:5]
	printSliceWithName("d", d)
}

func printSliceWithName(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlices() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeTest() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// index, value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	ppow := make([]int, 10)
	// only index
	for i := range ppow {
		ppow[i] = 1 << uint(i) // == 2**i
	}
	// only value
	for _, value := range ppow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	pointer()
	structTest()
	arrayTest()
	sliceTest()

	sliceDefaultIndexTest()
	makeSliceTest()
	sliceOfSlices()
	appendSlices()
	rangeTest()
}
