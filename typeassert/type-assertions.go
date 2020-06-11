package main

import "fmt"

// type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func testTypeSwitchs() {
	do(21)
	do("hello")
	do(true)
}
W
func testTypeAssert() {
	var i interface{} = "hello"

	// t := i.(T)
	s := i.(string)
	fmt.Println(s)

	// f = i.(float64) // panic
	// fmt.Println(f)

	// t, ok := i.(T)
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func main() {
	testTypeAssert()
	testTypeSwitchs()
}
