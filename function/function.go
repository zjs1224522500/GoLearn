package main

import "fmt"

// func func_name(arg1..) return_type
func add(x int, y int) int {
	return x + y
}

func addOmitType(x, y int) int {
	return x + y
}

// return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addOmitType(42, 13))

	// := declaration
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// 17 * 4 / 9 = 7(int)
	fmt.Println(split(17))

}
